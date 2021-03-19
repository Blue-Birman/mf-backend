package test

import (
	"testing"
	"time"

	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	"github.com/rvalessandro/mf-backend/modules/carts/services"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var cartDAOMock = &CartDAOMock{Mock: mock.Mock{}}
var cartService = services.CartService{DAO: cartDAOMock}

// Definisi objek mock
var time1, _ = time.Parse(time.RFC3339, "2020-12-08T00:00:00Z")
var time2, _ = time.Parse(time.RFC3339, "2020-12-08T00:00:00Z")
var product = productDomain.Product{
	ID:          1,
	CategoryID:  1,
	Name:        "White Clothes",
	Description: "White clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast",
	ImageURL:    "assets/product/set_1.png",
	PreviewURL:  "36805",
	Price:       1000000,
	CreatedAt:   time1,
	UpdatedAt:   time2,
}
var products = []productDomain.Product{product}

var now = time.Now()
var cart = domain.Cart{
	CustomerID: 1,
	ProductID:  1,
	CreatedAt:  now,
	UpdatedAt:  now,
}

func TestCartService_GetCartByCustomerIDNotFound(t *testing.T) {
	// Program mock data tidak ada
	cartDAOMock.Mock.On("Find", int64(0)).Return(nil)

	// Memanggil method Get Cart By Customer ID
	products, err := cartService.GetCartByCustomerID(int64(0))

	// Pengecekan bahwa products harus nil dan ada  error
	assert.Nil(t, products)
	assert.NotNil(t, err)

}

func TestCartService_GetCartByCustomerIDFound(t *testing.T) {
	// Program mock data ada
	cartDAOMock.Mock.On("Find", int64(1)).Return(products)

	// Memanggil method Get Cart By Customer ID
	result, err := cartService.GetCartByCustomerID(int64(1))

	// Pengecekan bahwa products sesuai
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, products[0].ID, (*result)[0].ID)
	assert.Equal(t, products[0].CategoryID, (*result)[0].CategoryID)
	assert.Equal(t, products[0].Name, (*result)[0].Name)

}

func TestCartService_CreateCartValid(t *testing.T) {
	cartParams := domain.CreateCartParams{
		CustomerID: 1,
		ProductID:  1,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Program mock data
	cartDAOMock.Mock.On("Create", cartParams).Return(cart)

	// Memanggil method CreateCart
	result, err := cartService.CreateCart(cartParams)

	// Pengecekan bahwa return sesuai dan tidak error
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, cart.ProductID, result.ProductID)
	assert.Equal(t, cart.CustomerID, result.CustomerID)

}

func TestCartService_CreateCartInvalid(t *testing.T) {
	var now = time.Now()
	cartParams := domain.CreateCartParams{
		CustomerID: -1,
		ProductID:  -1,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// Program mock data
	cartDAOMock.Mock.On("Create", cartParams).Return(nil)

	// Memanggil method CreateCart
	result, err := cartService.CreateCart(cartParams)

	// Pengecekan bahwa return sesuai dan tidak error
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCartService_DeleteCartValid(t *testing.T) {

	// Program mock data
	cartDAOMock.Mock.On("Get", int64(1), int64(1)).Return(cart)
	cartDAOMock.Mock.On("Delete", int64(1), int64(1)).Return(cart)

	// Memanggil method DeleteCart
	result, err := cartService.DeleteCart(int64(1), int64(1))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), result.ProductID)
	assert.Equal(t, int64(1), result.CustomerID)

}

func TestCartService_DeleteCartInvalid(t *testing.T) {

	// Program mock data
	cartDAOMock.Mock.On("Get", int64(-1), int64(-1)).Return(nil)
	cartDAOMock.Mock.On("Delete", int64(-1), int64(-1)).Return(nil)

	// Memanggil method DeleteCart
	_, err := cartService.DeleteCart(int64(-1), int64(-1))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.NotNil(t, err)

}
