package test

import (
	"testing"

	"github.com/rvalessandro/mf-backend/modules/carts/services"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cartDAO = &CartDAOMock{Mock: mock.Mock{}}
var cartService = services.CartService{DAO: cartDAO}

func TestCartService_GetNotFound(t *testing.T) {
	// Program mock data tidak ada
	cartDAO.Mock.On("Find", int64(1)).Return(nil)

	// Memanggil method Get Cart By Customer ID
	products, err := cartService.GetCartByCustomerID(int64(1))

	// Pengecekan bahwa products harus nil dan ada  error
	assert.Nil(t, products)
	assert.NotNil(t, err)

}

func TestCartService_GetFound(t *testing.T) {
	products := []productDomain.Product{
		{
			ID:          2,
			CategoryID:  1,
			Name:        "Sweater",
			Description: "Comfy Sweater",
			ImageURL:    "folder/",
			PreviewURL:  "google.com",
			Price:       100000,
		},
	}

	// Program mock data ada
	cartDAO.Mock.On("Find", int64(2)).Return(products)

	// Memanggil method Get Cart By Customer ID
	result, err := cartService.GetCartByCustomerID(int64(2))

	// Pengecekan bahwa products harus nil dan ada  error
	assert.NotNil(t, result)
	assert.Nil(t, err)

}
