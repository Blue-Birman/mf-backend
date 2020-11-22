package presentation

import (
	"github.com/gin-gonic/gin"
	customerController "github.com/rvalessandro/mf-backend/modules/products/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/product", customerController.Find)
	router.GET("/product/:customer_id", customerController.Get)
	router.POST("/product/", customerController.Create)
	router.PUT("/product/:customer_id", customerController.Update)
	router.DELETE("/product/:customer_id", customerController.Delete)
}
