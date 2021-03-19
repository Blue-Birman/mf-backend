package services

import (
	"time"

	// "github.com/rvalessandro/mf-backend/modules/carts/data"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

type CartDAOInterface interface {
	Find(customerID int64) (*[]productDomain.Product, *errors.RestErr)
	Get(customerID int64, productID int64) (*domain.Cart, *errors.RestErr)
	Create(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr)
	Delete(customerID int64, productID int64) *errors.RestErr
}

type CartService struct {
	DAO CartDAOInterface
}

func (service CartService) GetCartByCustomerID(customerID int64) (*[]productDomain.Product, *errors.RestErr) {
	carts, err := service.DAO.Find(customerID)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (service CartService) CreateCart(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr) { // SaveToCart
	currentTime := time.Now()
	cartParam.CreatedAt = currentTime
	cartParam.UpdatedAt = currentTime

	cart, err := service.DAO.Create(cartParam) // cart
	if err != nil {
		return nil, err
	}

	return cart, nil // cart, r = data.get()
}

func (service CartService) DeleteCart(customerID int64, productID int64) (*domain.Cart, *errors.RestErr) {
	cart, err := service.DAO.Get(customerID, productID)
	if err != nil {
		return nil, err
	}

	err = service.DAO.Delete(customerID, productID)
	if err != nil {
		return nil, err
	}

	return cart, err
}
