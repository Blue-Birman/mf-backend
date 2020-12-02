package presentation

import (
	"github.com/gin-gonic/gin"
	transactionController "github.com/rvalessandro/mf-backend/modules/transactions/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/transactions", transactionController.Find)
	router.GET("/transactions/:transaction_id", transactionController.Get)
	router.POST("/transactions", transactionController.Create)
	router.PUT("/transactions/:transaction_id", transactionController.Update)
	router.DELETE("/transactions/:transaction_id", transactionController.Delete)
}
