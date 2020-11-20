package presentation

import (
	"github.com/gin-gonic/gin"
	productController "github.com/rvalessandro/mf-backend/modules/products/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/products/:product_id", productController.Get)
}
