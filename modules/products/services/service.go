package services

import (
	"github.com/rvalessandro/mf-backend/modules/products/data"
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetProduct(productID int64) (*domain.Product, *errors.RestErr) {
	product, err := data.Get()
	if err != nil {
		return nil, err
	}
	return product, nil
}
