package presentation

import (
	"github.com/gin-gonic/gin"
	transactionController "github.com/rvalessandro/mf-backend/modules/transactions/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/transaction", transactionController.Find)
	router.GET("/transaction/:transaction_id", transactionController.Get)
	router.POST("/transaction/", transactionController.Create)
	router.PUT("/transaction/:transaction_id", transactionController.Update)
	router.DELETE("/transaction/:transaction_id", transactionController.Delete)
}
