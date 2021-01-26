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

func Update(c *gin.Context) {
	customerID, err := parser.ParseID(c.Param("customer_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var customerParam domain.CreateCustomerParams
	bindErr := c.ShouldBindJSON(&customerParam)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	customer, updateErr := services.UpdateCustomer(customerID, customerParam)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func Delete(c *gin.Context) {
	customerID, err := parser.ParseID(c.Param("customer_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var customer *domain.Customer
	customer, err = services.DeleteCustomer(customerID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}
