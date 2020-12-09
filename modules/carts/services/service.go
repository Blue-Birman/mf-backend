package services

import (
	"fmt"
	"time"

	"github.com/rvalessandro/mf-backend/modules/carts/data"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetCartByCustomerID(customerID int64) ([]productDomain.Product, *errors.RestErr) {
	carts, err := data.Find(customerID)
	fmt.Println("masuk")
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func CreateCart(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr) {
	currentTime := time.Now()
	cartParam.CreatedAt = currentTime
	cartParam.UpdatedAt = currentTime

	cart, err := data.Create(cartParam)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func DeleteCart(customerID int64, productID int64) (*domain.Cart, *errors.RestErr) {
	cart, err := data.Get(customerID, productID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(customerID, productID)
	if err != nil {
		return nil, err
	}

	return cart, err
}
