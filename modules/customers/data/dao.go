package data

import (
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/modules/customers/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetCustomer = "SELECT id, name, email, address, created_at, updated_at FROM customers WHERE id=?"
)

func Get() (*domain.Customer, *errors.RestErr) {
	customer := domain.Customer{}
	stmt, err := mysql.Client.Prepare(queryGetCustomer)
	if err != nil {
		return nil, errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(customer.ID)
	if getErr := result.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Address, &customer.CreatedAt, &customer.UpdatedAt); getErr != nil {
		return nil, mysql_util.ParseError(getErr)
	}

	return &customer, nil
}
