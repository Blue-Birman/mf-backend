package customers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/customers/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	customerID, idErr := parser.ParseID(c.Param("customer_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	customer, saveErr := services.GetCustomer(customerID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, customer)
}
