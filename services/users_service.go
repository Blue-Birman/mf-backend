package services

import (
	"github.com/rvalessandro/mf-backend/domain/customers"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetUser(customerID int64) (*customers.Customer, *errors.RestErr) {
	customer := &customers.Customer{ID: customerID}
	if err := customer.Get(); err != nil {
		return nil, err
	}
	return customer, nil
}
