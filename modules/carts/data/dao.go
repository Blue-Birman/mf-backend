package data

import (
	"fmt"

	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetCart              = `SELECT customer_id, product_id, created_at, updated_at from carts where customer_id=? and product_id=?;`
	queryGetCartsByCustomerID = `SELECT customer_id, product_id, created_at, updated_at from carts where customer_id=?;`
	queryCreateCart           = `
		INSERT INTO carts (customer_id, product_id, created_at, updated_at)
		VALUES (
			?, ?, ?, ?
		)`
	queryDeleteCart = `DELETE FROM carts WHERE customer_id=$1 AND product_id=$1`
)

// func Find(customerID int64) (domain.Cart, *errors.RestErr) {
// 	cart := domain.Cart{}
// 	stmt, err := mysql.Client.Prepare(queryGetCartsByCustomerID)
// 	if err != nil {
// 		return nil, errors.NewErrInternalServer(err.Error())
// 	}
// 	defer stmt.Close()

// 	result := stmt.QueryRow(customerID)
// 	fmt.Println(result)
// 	err = result.Scan(&cart.CustomerID, &cart.ProductID, &cart.CreatedAt, &cart.UpdatedAt)
// 	if err != nil {
// 		return nil, mysql_util.ParseError(err)
// 	}

// 	return &cart, nil
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
		var cart domain.Cart
		err = rows.Scan(&cart.CustomerID, &cart.ProductID, &cart.CreatedAt, &cart.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, cart)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No carts found"))
	}

	return results, nil
}

func Get(customerID int64, productID int64) (*domain.Cart, *errors.RestErr) {
	cart := domain.Cart{}
	stmt, err := mysql.Client.Prepare(queryGetCart)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(customerID, productID)
	err = result.Scan(&cart.CustomerID, &cart.ProductID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &cart, nil
}

func Create(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateCart)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	stmt.Exec(&cartParam.CustomerID, &cartParam.ProductID, &cartParam.CreatedAt, &cartParam.UpdatedAt)

	// queryRes, err := stmt.Exec(&cartParam.CustomerID, &cartParam.ProductID, &cartParam.CreatedAt, &cartParam.UpdatedAt)
	// if err != nil {
	// 	return nil, mysql_util.ParseError(err)
	// }

	// newID, err := queryRes.LastInsertId()

	return Get(cartParam.CustomerID, cartParam.ProductID)
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
