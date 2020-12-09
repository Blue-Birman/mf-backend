package services

import (
	"github.com/rvalessandro/mf-backend/modules/categories/data"
	"github.com/rvalessandro/mf-backend/modules/categories/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"time"
)

func FindCategory() ([]domain.Category, *errors.RestErr) {
	categories, err := data.Find()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func GetCategory(customerID int64) (*domain.Category, *errors.RestErr) {
	customer, err := data.Get(customerID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func CreateCategory(customerParam domain.CreateCategoryParams) (*domain.Category, *errors.RestErr) {
	customer, err := data.Create(customerParam)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func UpdateCategory(id int64, customerParam domain.CreateCategoryParams) (*domain.Category, *errors.RestErr) {
	customerParam.UpdatedAt = time.Now()

	customer, err := data.Update(id, customerParam)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func DeleteCategory(customerID int64) (*domain.Category, *errors.RestErr) {
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
