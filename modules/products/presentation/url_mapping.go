package presentation

import (
	"github.com/gin-gonic/gin"
	productController "github.com/rvalessandro/mf-backend/modules/products/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/product", productController.Find)
	router.GET("/product/:product_id", productController.Get)
	router.POST("/product/", productController.Create)
	router.PUT("/product/:product_id", productController.Update)
	router.DELETE("/product/:product_id", productController.Delete)
}
