package services

import (
	"time"

	"github.com/rvalessandro/mf-backend/modules/carts/data"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetCartByCustomerID(customerID int64) (*domain.Cart, *errors.RestErr) {
	cart, err := data.Get(customerID)
	if err != nil {
		return nil, err
	}

	return cart, nil
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

func DeleteCart(cartID int64) (*domain.Cart, *errors.RestErr) {
	cart, err := data.Get(cartID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(cartID)
	if err != nil {
		return nil, err
	}

	return cart, err
}
