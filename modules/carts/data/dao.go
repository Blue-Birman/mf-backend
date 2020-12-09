package data

import (
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetCartsByCustomerID = `SELECT customer_id, product_id, created_at, updated_at from categories where customer_id=?;`
	queryCreateCart           = `
		INSERT INTO categories (customer_id, product_id, created_at, updated_at)
		VALUES (
			?, ?, ?, ?
		)`
	queryDeleteCart = `DELETE FROM carts WHERE customer_id=$1 AND product_id=$1`
)

func Get(cartID int64) (*domain.Cart, *errors.RestErr) {
	cart := domain.Cart{}
	stmt, err := mysql.Client.Prepare(queryGetCartsByCustomerID)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(cartID)
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

	queryRes, err := stmt.Exec((&cart.CustomerID, &cart.ProductID, &cart.CreatedAt, &cart.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	newID, err := queryRes.LastInsertId()

	return Get(newID)
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
