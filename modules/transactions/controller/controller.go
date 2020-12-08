package transactions

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/transactions/domain"
	"github.com/rvalessandro/mf-backend/modules/transactions/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
	"net/http"
)

func Find(c *gin.Context) {
	products, err := services.FindTransaction()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func Get(c *gin.Context) {
	productID, idErr := parser.ParseID(c.Param("transaction_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, err := services.GetTransaction(productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Create(c *gin.Context) {
	var transactionParam domain.CreateTransactionParams
	err := c.ShouldBindJSON(&transactionParam)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newTransaction, createErr := services.CreateTransaction(transactionParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}

func Update(c *gin.Context) {
	productID, err := parser.ParseID(c.Param("product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var productParam domain.CreateTransactionParams
	bindErr := c.ShouldBindJSON(&productParam)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	product, updateErr := services.UpdateTransaction(productID, productParam)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Delete(c *gin.Context) {
	productID, err := parser.ParseID(c.Param("product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var product *domain.Transaction
	product, err = services.DeleteTransaction(productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}
