package wishlists

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/modules/wishlists/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	customerID, idErr := parser.ParseID(c.Param("customer_id"))
	fmt.Println("Geet")
	fmt.Println(customerID)
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	products, err := services.GetWishlistByCustomerID(customerID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func Create(c *gin.Context) {
	var wishlistParam = domain.NewWishlist()
	fmt.Println(wishlistParam)
	err := c.ShouldBindJSON(&wishlistParam)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newWishlist, createErr := services.CreateWishlist(*wishlistParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newWishlist)
}

func Delete(c *gin.Context) {
	customerID, err := parser.ParseID(c.Param("customer_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	productID, err := parser.ParseID(c.Param("product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var product *domain.Wishlist
	product, err = services.DeleteWishlist(customerID, productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}
