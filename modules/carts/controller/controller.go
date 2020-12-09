package carts

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/carts/domain"
	"github.com/rvalessandro/mf-backend/modules/carts/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	customerID, idErr := parser.ParseID(c.Param("customer_id"))
	fmt.Println("Geet")
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	products, err := services.GetCartByCustomerID(customerID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func Create(c *gin.Context) {
	var cartParam domain.CreateCartParams
	err := c.ShouldBindJSON(&cartParam)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newCart, createErr := services.CreateCart(cartParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newCart)
}

func Delete(c *gin.Context) {
	productID, err := parser.ParseID(c.Param("customer_id", "product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var product *domain.Cart
	product, err = services.DeleteCart(productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}
