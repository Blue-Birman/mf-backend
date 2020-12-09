package data

import (
	"fmt"

	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetCart              = `SELECT customer_id, product_id, created_at, updated_at from wishlists where customer_id=? and product_id=?;`
	queryGetCartsByCustomerID = `SELECT customer_id, product_id, created_at, updated_at from wishlists where customer_id=?;`
	queryCreateCart           = `
		INSERT INTO wishlists (customer_id, product_id, created_at, updated_at)
		VALUES (
			?, ?, ?, ?
		)`
	queryDeleteCart = `DELETE FROM wishlists WHERE customer_id=? AND product_id=?`
)

// func Find(customerID int64) (domain.Cart, *errors.RestErr) {
// 	wishlist := domain.Cart{}
// 	stmt, err := mysql.Client.Prepare(queryGetCartsByCustomerID)
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

func Find(customerID int64) ([]domain.Cart, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryGetCartsByCustomerID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(customerID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Cart, 0)
	for rows.Next() {
		var wishlist domain.Cart
		err = rows.Scan(&wishlist.CustomerID, &wishlist.ProductID, &wishlist.CreatedAt, &wishlist.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, wishlist)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No wishlists found"))
	}

	return results, nil
}

func Get(customerID int64, productID int64) (*domain.Cart, *errors.RestErr) {
	wishlist := domain.Cart{}
	stmt, err := mysql.Client.Prepare(queryGetCart)
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

func Create(wishlistParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateCart)
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
	stmt, err := mysql.Client.Prepare(queryDeleteCart)
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
