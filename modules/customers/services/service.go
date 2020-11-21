package services

import (
	"github.com/rvalessandro/mf-backend/modules/customers/data"
	"github.com/rvalessandro/mf-backend/modules/customers/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"time"
)

func FindCustomer() ([]domain.Customer, *errors.RestErr) {
	customers, err := data.Find()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func GetCustomer(customerID int64) (*domain.Customer, *errors.RestErr) {
	customer, err := data.Get(customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateCustomer(customerParam domain.CreateCustomerParams) (*domain.Customer, *errors.RestErr) {
	currentTime := time.Now()
	customerParam.CreatedAt = currentTime
	customerParam.UpdatedAt = currentTime

	customer, err := data.Create(customerParam)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func DeleteCustomer(customerID int64) (*domain.Customer, *errors.RestErr) {
	customer, err := data.Get(customerID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(customerID)
	if err != nil {
		return nil, err
	}

	return customer, err
}
