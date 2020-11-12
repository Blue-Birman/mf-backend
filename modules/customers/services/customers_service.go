package services

import (
	"github.com/rvalessandro/mf-backend/modules/customers/data"
	"github.com/rvalessandro/mf-backend/modules/customers/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetCustomer(customerID int64) (*domain.Customer, *errors.RestErr) {
	customer, err := data.Get()
	if err != nil {
		return nil, err
	}
	return customer, nil
}
