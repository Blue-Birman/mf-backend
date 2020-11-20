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
	queryGetCustomer    = `SELECT id, name, email, address, created_at, updated_at from customers where id=$1;`
	queryCreateCustomer = `
		INSERT INTO customers (name, email, password, address, created_at)
		VALUES (
			$1, $2, $3, $4, $5
		)
		RETURNING id, name, email, password, address, created_at, updated_at`
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
