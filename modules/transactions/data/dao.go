package data

import (
	"fmt"
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	productDAO "github.com/rvalessandro/mf-backend/modules/products/data"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/modules/transactions/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryFindTransaction = `SELECT id, customer_id, date, created_at, updated_at FROM transactions;`
	queryGetTransaction  = `
		SELECT id, customer_id, date, created_at, updated_at 
		FROM transactions 
		WHERE id=? 
		LIMIT 1;
	`
	queryCreateTransaction = `
		INSERT INTO transactions (customer_id, date, created_at, updated_at)
		VALUES (?, ?, ?, ?);
	`
	queryCreateTransactionProducts = `
		INSERT INTO transaction_products (transaction_id, product_id, qty, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?);
	`
	queryUpdateTransaction = `
		UPDATE transactions
		SET 
		    customer_id=?,
		    date=?,
		    updated_at=?
		WHERE id=?
	`
	queryDeleteTransaction = `DELETE FROM transactions WHERE id=?`

	queryFindTransactionProductIDs = `SELECT product_id, qty FROM transaction_products WHERE transaction_id=?`
)

func Find() ([]domain.Transaction, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindTransaction)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Transaction, 0)
	for rows.Next() {
		var transaction domain.Transaction
		err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, transaction)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No transactions found"))
	}

	for i, transaction := range results {
		products, productsErr := FindTransactionProducts(transaction.ID)
		if productsErr != nil {
			return nil, productsErr
		}

		results[i].Products = products
	}

	return results, nil
}

func Get(id int64) (*domain.Transaction, *errors.RestErr) {
	transaction := domain.Transaction{}
	stmt, err := mysql.Client.Prepare(queryGetTransaction)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(id)
	err = result.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	products, productsErr := FindTransactionProducts(id)
	if productsErr != nil {
		return nil, productsErr
	}

	transaction.Products = products

	return &transaction, nil
}

func Create(TransactionParam domain.CreateTransactionParams) (*domain.Transaction, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateTransaction)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	queryRes, err := stmt.Exec(TransactionParam.CustomerID, TransactionParam.Date, TransactionParam.CreatedAt, TransactionParam.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}
	newID, err := queryRes.LastInsertId()

	/**
	 * Insert Transaction Products
	 */
	prodStmt, prodErr := mysql.Client.Prepare(queryCreateTransactionProducts)
	if prodErr != nil {
		return nil, errors.NewErrInternalServer(prodErr.Error())
	}
	defer prodStmt.Close()

	for _, product := range TransactionParam.Products {
		product.TransactionID = newID
		_, prodErr := prodStmt.Exec(product.TransactionID, product.ProductID, product.Qty, TransactionParam.CreatedAt, TransactionParam.UpdatedAt)
		if prodErr != nil {
			return nil, errors.NewErrInternalServer(prodErr.Error())
		}
	}

	return Get(newID)
}

func Update(id int64, TransactionParam domain.CreateTransactionParams) (*domain.Transaction, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryUpdateTransaction)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(TransactionParam.Date, TransactionParam.UpdatedAt, id)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}

	return Get(id)
}

func Delete(id int64) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteTransaction)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}

	return nil
}

/**
 *
 */
func FindTransactionProducts(transactionID int64) ([]productDomain.TransactionProduct, *errors.RestErr) {
	findProductIDsStmt, err := mysql.Client.Prepare(queryFindTransactionProductIDs)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer findProductIDsStmt.Close()

	rows, err := findProductIDsStmt.Query(transactionID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}

	transactionProducts := make([]productDomain.TransactionProduct, 0)
	for rows.Next() {
		var transactionProduct productDomain.TransactionProduct
		err := rows.Scan(&transactionProduct.ID, &transactionProduct.Qty)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		transactionProducts = append(transactionProducts, transactionProduct)
	}

	products := make([]productDomain.TransactionProduct, 0)
	for _, transactionProduct := range transactionProducts {
		product, err := productDAO.Get(transactionProduct.ID)
		transactionProduct.Name = product.Name
		transactionProduct.Description = product.Name
		transactionProduct.ImageURL = product.Name
		transactionProduct.Price = product.Price
		transactionProduct.Total = transactionProduct.Price * int64(transactionProduct.Qty)
		if err != nil {
			return nil, err
		}
		products = append(products, transactionProduct)
	}

	return products, nil
}
