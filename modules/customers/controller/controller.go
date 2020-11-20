package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/customers/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Find(c *gin.Context) {
	customers, err := services.FindCustomer()
	if err != nil {
		c.JSON(err.Status, err)
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

	c.JSON(http.StatusCreated, customer)
}
