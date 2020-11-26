package data

import (
	"fmt"
	"github.com/rvalessandro/mf-backend/datasources/mysql"
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
		VALUES (?,?,?,?);
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
		err := rows.Scan(&transaction.ID, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, transaction)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No transactions found"))
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
	err = result.Scan(&transaction.ID, &transaction.Date, &transaction.CreatedAt, &transaction.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &transaction, nil
}

func Create(TransactionParam domain.CreateTransactionParams) (*domain.Transaction, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateTransaction)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	queryRes, err := stmt.Exec(TransactionParam.Date, TransactionParam.CreatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	newID, err := queryRes.LastInsertId()

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
