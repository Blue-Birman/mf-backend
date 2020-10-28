package customers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/go-bookstore_users-api/services"
	"github.com/rvalessandro/go-bookstore_users-api/utils/errors"
)

func getCustomerID(customerIDParam string) (int64, *errors.RestErr) {
	customerID, customerErr := strconv.ParseInt(customerIDParam, 10, 64)
	if customerErr != nil {
		return 0, errors.NewErrBadRequest("invalid user id")
	}

	return customerID, nil
}

func Get(c *gin.Context) {
	userID, idErr := getCustomerID(c.Param("customer_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, saveErr := services.GetUser(userID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, user)
}
