package wishlists

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/wishlists/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	wishlistID, idErr := parser.ParseID(c.Param("wishlist_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	wishlist, saveErr := services.GetWishlist(wishlistID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, wishlist)
}
