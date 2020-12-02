package presentation

import (
	"github.com/gin-gonic/gin"
	productController "github.com/rvalessandro/mf-backend/modules/products/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/products", productController.Find)
	router.GET("/products/:product_id", productController.Get)
	router.POST("/products", productController.Create)
	router.PUT("/products/:product_id", productController.Update)
	router.DELETE("/products/:product_id", productController.Delete)
}
