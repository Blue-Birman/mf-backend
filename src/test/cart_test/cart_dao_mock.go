package test

import (
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/stretchr/testify/mock"
)

type CartDAOInterface interface {
	Find(customerID int64) (*[]productDomain.Product, *errors.RestErr)
	Get(customerID int64, productID int64) (*domain.Cart, *errors.RestErr)
	Create(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr)
	Delete(customerID int64, productID int64) *errors.RestErr
}

type CartDAOMock struct {
	Mock mock.Mock
}

func (dao *CartDAOMock) Find(customerID int64) (*[]productDomain.Product, *errors.RestErr) {
	arguments := dao.Mock.Called(customerID)
	if arguments.Get(0) == nil {
		return nil, errors.NewErrInternalServer("User not Found!")
	} else {
		products := arguments.Get(0).([]productDomain.Product)
		return &products, nil
	}
}

func (dao *CartDAOMock) Get(customerID int64, productID int64) (*domain.Cart, *errors.RestErr) {
	return nil, nil
}

func (dao *CartDAOMock) Create(cartParam domain.CreateCartParams) (*domain.Cart, *errors.RestErr) {
	return nil, nil
}

func (dao *CartDAOMock) Delete(customerID int64, productID int64) *errors.RestErr {
	return nil
}
