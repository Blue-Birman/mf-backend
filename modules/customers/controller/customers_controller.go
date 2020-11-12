package customers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/services"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func getCustomerID(customerIDParam string) (int64, *errors.RestErr) {
	customerID, customerErr := strconv.ParseInt(customerIDParam, 10, 64)
	if customerErr != nil {
		return 0, errors.NewErrBadRequest("invalid customer id")
	}

	return customerID, nil
}

func Get(c *gin.Context) {
	customerID, idErr := getCustomerID(c.Param("customer_id"))
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
