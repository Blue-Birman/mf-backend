package data

import (
	"fmt"
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/customers/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryFindCustomer   = `SELECT id, name, email, address, created_at, updated_at FROM customers`
	queryGetCustomer    = `SELECT id, name, email, address, created_at, updated_at from customers where id=?;`
	queryCreateCustomer = `
		INSERT INTO customers (name, email, password, address, created_at)
		VALUES (
			?, ?, ?, ?, ?
		)`
	queryUpdateCustomer = `UPDATE customers SET name=?, email=?, password=?, address=?, updated_at=? WHERE id=?`
	queryDeleteCustomer = `DELETE FROM customers WHERE id=$1`
)

func Find() ([]domain.Customer, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryFindCustomer)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer rows.Close()

	results := make([]domain.Customer, 0)
	for rows.Next() {
		var customer domain.Customer
		err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
		if err != nil {
			return nil, errors.NewErrInternalServer(err.Error())
		}
		results = append(results, customer)
	}

	if len(results) == 0 {
		return nil, errors.NewErrNotFound(fmt.Sprintf("No users found"))
	}

	return results, nil
}

func Get(customerID int64) (*domain.Customer, *errors.RestErr) {
	customer := domain.Customer{}
	stmt, err := mysql.Client.Prepare(queryGetCustomer)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(customerID)
	err = result.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	return &customer, nil
}

func Create(customerParam domain.CreateCustomerParams) (*domain.Customer, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryCreateCustomer)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	queryRes, err := stmt.Exec(customerParam.Name, customerParam.Email, customerParam.Password, customerParam.Address, customerParam.CreatedAt)
	if err != nil {
		return nil, mysql_util.ParseError(err)
	}

	newID, err := queryRes.LastInsertId()

	return Get(newID)
}

func Update(id int64, customerParam domain.CreateCustomerParams) (*domain.Customer, *errors.RestErr) {
	stmt, err := mysql.Client.Prepare(queryUpdateCustomer)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(customerParam.Name, customerParam.Email, customerParam.Password, customerParam.Address, customerParam.UpdatedAt, id)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}

	return Get(id)
}

func Delete(customerID int64) *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryDeleteCustomer)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(customerID)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}

	return nil
}
