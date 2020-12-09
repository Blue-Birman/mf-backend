package data

import (
	"fmt"

	"github.com/rvalessandro/mf-backend/datasources/mysql"
	productDAO "github.com/rvalessandro/mf-backend/modules/products/data"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetWishlist              = `SELECT customer_id, product_id, created_at, updated_at from wishlists where customer_id=? and product_id=?;`
	queryGetWishlistsByCustomerID = `SELECT customer_id, product_id, created_at, updated_at from wishlists where customer_id=?;`
	queryCreateWishlist           = `
		INSERT INTO wishlists (customer_id, product_id, created_at, updated_at)
		VALUES (
			?, ?, ?, ?
		)`
	queryDeleteWishlist = `DELETE FROM wishlists WHERE customer_id=? AND product_id=?`
)

// func Find(customerID int64) (domain.Wishlist, *errors.RestErr) {
// 	wishlist := domain.Wishlist{}
// 	stmt, err := mysql.Client.Prepare(queryGetWishlistsByCustomerID)
// 	if err != nil {
// 		return nil, errors.NewErrInternalServer(err.Error())
// 	}
// 	defer stmt.Close()

// 	result := stmt.QueryRow(customerID)
// 	fmt.Println(result)
// 	err = result.Scan(&wishlist.CustomerID, &wishlist.ProductID, &wishlist.CreatedAt, &wishlist.UpdatedAt)
// 	if err != nil {
// 		return nil, mysql_util.ParseError(err)
// 	}

// 	return &wishlist, nil
// }

func Find(customerID int64) ([]productDomain.Product, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryGetWishlistsByCustomerID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Wishlist, 0)
	for rows.Next() {
		var wishlist domain.Wishlist
		err = rows.Scan(&wishlist.CustomerID, &wishlist.ProductID, &wishlist.CreatedAt, &wishlist.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, wishlist)
	}

	products := make([]productDomain.Product, 0)
	for _, transactionProduct := range results {
		product, err := productDAO.Get(transactionProduct.ProductID)
		fmt.Println(transactionProduct.ProductID)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No wishlists found"))
	}

	return products, nil
}

func Get(customerID int64, productID int64) (*domain.Wishlist, *errors.RestErr) {
	wishlist := domain.Wishlist{}
	stmt, err := mysql.Client.Prepare(queryGetWishlist)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(customerID, productID)
	err = result.Scan(&wishlist.CustomerID, &wishlist.ProductID, &wishlist.CreatedAt, &wishlist.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &wishlist, nil
}

func Create(wishlistParam domain.CreateWishlistParams) (*domain.Wishlist, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateWishlist)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	stmt.Exec(&wishlistParam.CustomerID, &wishlistParam.ProductID, &wishlistParam.CreatedAt, &wishlistParam.UpdatedAt)

	// queryRes, err := stmt.Exec(&wishlistParam.CustomerID, &wishlistParam.ProductID, &wishlistParam.CreatedAt, &wishlistParam.UpdatedAt)
	// if err != nil {
	// 	return nil, mysql_util.ParseError(err)
	// }

	// newID, err := queryRes.LastInsertId()

	return Get(wishlistParam.CustomerID, wishlistParam.ProductID)
}

func Delete(customerID int64, productID int64) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteWishlist)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(customerID, productID)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}

	return nil
}
