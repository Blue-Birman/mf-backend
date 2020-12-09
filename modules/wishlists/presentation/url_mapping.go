package presentation

import (
	"github.com/gin-gonic/gin"
	wishlistController "github.com/rvalessandro/mf-backend/modules/wishlists/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/wishlists/:customer_id", wishlistController.Get)
	router.POST("/wishlists", wishlistController.Create)
	router.DELETE("/wishlists/:customer_id/:product_id", wishlistController.Delete)
}
