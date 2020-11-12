package transactions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/transactions/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	transactionID, idErr := parser.ParseID(c.Param("transaction_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	transaction, saveErr := services.GetTransaction(transactionID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, transaction)
}
