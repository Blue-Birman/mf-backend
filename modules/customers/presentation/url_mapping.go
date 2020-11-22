package presentation

import (
	"github.com/gin-gonic/gin"
	customerController "github.com/rvalessandro/mf-backend/modules/customers/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/customers", customerController.Find)
	router.GET("/customers/:customer_id", customerController.Get)
	router.POST("/customers/", customerController.Create)
	router.PUT("/customers/:customer_id", customerController.Update)
	router.DELETE("/customers/:customer_id", customerController.Delete)
}
