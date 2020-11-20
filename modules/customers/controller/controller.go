package customers

import (
	"github.com/rvalessandro/mf-backend/modules/customers/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/customers/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Find(c *gin.Context) {
	customers, err := services.FindCustomer()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func Get(c *gin.Context) {
	customerID, idErr := parser.ParseID(c.Param("customer_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	customer, err := services.GetCustomer(customerID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func Create(c *gin.Context) {
	var customerParam domain.CreateCustomerParams
	err := c.ShouldBindJSON(&customerParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newCustomer, createErr := services.CreateCustomer(customerParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newCustomer)
}
