package presentation

import (
	"github.com/gin-gonic/gin"
	wishlistController "github.com/rvalessandro/mf-backend/modules/wishlists/controller"
)

func MapURLs(router gin.Engine) {
	router.GET("/wishlists/:wishlist_id", wishlistController.Get)
}
