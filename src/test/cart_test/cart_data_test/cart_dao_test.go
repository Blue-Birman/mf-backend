package testdao

import (
	"fmt"
	"testing"
	"time"

	"github.com/rvalessandro/mf-backend/modules/carts/data"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/stretchr/testify/assert"
)

var cartDAO = &data.CartDAO{}

func TestCartDAO_FindFound(t *testing.T) {
	time1, _ := time.Parse(time.RFC3339, "2020-12-08T00:00:00Z")
	time2, _ := time.Parse(time.RFC3339, "2020-12-08T00:00:00Z")
	products := []productDomain.Product{
		{
			ID:          1,
			CategoryID:  1,
			Name:        "Blue Clothes",
			Description: "Blue clothes with elastic fiber, appropiate for hard sporting, will absorb sweat fast",
			ImageURL:    "assets/product/set_1.png",
			PreviewURL:  "36805",
			Price:       1000000,
			CreatedAt:   time1,
			UpdatedAt:   time2,
		},
	}

	// Memanggil method Find
	result, err := cartDAO.Find(int64(6))

	// Pengecekan bahwa products tidak nil dan tidak error
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, products[0].ID, (*result)[0].ID)
	assert.Equal(t, products[0].CategoryID, (*result)[0].CategoryID)
	assert.Equal(t, products[0].Name, (*result)[0].Name)

}

func TestCartDAO_FindNotFound(t *testing.T) {
	// Memanggil method Find
	result, err := cartDAO.Find(int64(6))

	// Pengecekan bahwa products harus nil dan ada error
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCartDAO_GetFound(t *testing.T) {
	customerID := int64(6)
	productID := int64(1)

	// Memanggil method Get
	result, err := cartDAO.Get(int64(6), int64(1))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, customerID, (*result).CustomerID)
	assert.Equal(t, productID, (*result).ProductID)

}

func TestCartDAO_GetNotFound(t *testing.T) {
	// Memanggil method Get
	result, err := cartDAO.Get(int64(6), int64(1))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestCartDAO_CreateValid(t *testing.T) {
	cartParams := domain.CreateCartParams{
		CustomerID: 6,
		ProductID:  2,
	}
	// Memanggil method Create
	result, err := cartDAO.Create(cartParams)

	// Pengecekan bahwa return sesuai dan tidak error
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, cartParams.CustomerID, result.CustomerID)
	assert.Equal(t, cartParams.ProductID, result.ProductID)
}

func TestCartDAO_CreateInvalid(t *testing.T) {
	cartParams := domain.CreateCartParams{
		CustomerID: -1,
		ProductID:  -1,
	}
	// Memanggil method Create
	_, err := cartDAO.Create(cartParams)
	fmt.Println(err)

	// Pengecekan bahwa return error
	assert.NotNil(t, err)
}

func TestCartDAO_DeleteValid(t *testing.T) {
	// Memanggil method Delete
	err := cartDAO.Delete(int64(6), int64(2))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.Nil(t, err)
}

func TestCartDAO_DeleteInvalid(t *testing.T) {
	// Memanggil method Delete
	err := cartDAO.Delete(int64(0), int64(0))

	// Pengecekan bahwa return sesuai dan tidak error
	assert.Nil(t, err)
}
