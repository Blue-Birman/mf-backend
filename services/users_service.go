package services

import (
	"github.com/rvalessandro/go-bookstore_users-api/domain/customers"
	"github.com/rvalessandro/go-bookstore_users-api/utils/errors"
)

func GetUser(customerID int64) (*customers.Customer, *errors.RestErr) {
	customer := &customers.Customer{ID: customerID}
	if err := customer.Get(); err != nil {
		return nil, err
	}
	return customer, nil
}
