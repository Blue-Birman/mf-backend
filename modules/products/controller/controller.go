package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/products/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Get(c *gin.Context) {
	productID, idErr := parser.ParseID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, saveErr := services.GetProduct(productID)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, product)
}
