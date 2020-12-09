package presentation

import (
	"github.com/gin-gonic/gin"
	cartController "github.com/rvalessandro/mf-backend/modules/carts/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/carts/:customer_id", cartController.Get)
	router.POST("/carts", cartController.Create)
	router.DELETE("/carts/:customer_id/:product_id", cartController.Delete)
}
