package presentation

import (
	"github.com/gin-gonic/gin"
	transactionController "github.com/rvalessandro/mf-backend/modules/transactions/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/transactions/:transaction_id", transactionController.Get)
}
