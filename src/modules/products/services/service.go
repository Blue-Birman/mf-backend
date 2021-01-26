package services

import (
	"time"

	"github.com/rvalessandro/mf-backend/modules/products/data"
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func FindProduct() ([]domain.Product, *errors.RestErr) {
	products, err := data.Find()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductsByCategory(categoryID int64) ([]domain.Product, *errors.RestErr) {
	products, err := data.GetByCategory(categoryID)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProduct(productID int64) (*domain.Product, *errors.RestErr) {
	product, err := data.Get(productID)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func CreateProduct(productParam domain.CreateProductParams) (*domain.Product, *errors.RestErr) {
	currentTime := time.Now()
	productParam.CreatedAt = currentTime
	productParam.UpdatedAt = currentTime

	product, err := data.Create(productParam)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProduct(id int64, productParam domain.CreateProductParams) (*domain.Product, *errors.RestErr) {
	productParam.UpdatedAt = time.Now()

	product, err := data.Update(id, productParam)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func DeleteProduct(productID int64) (*domain.Product, *errors.RestErr) {
	product, err := data.Get(productID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(productID)
	if err != nil {
		return nil, err
	}

	return product, err
}
