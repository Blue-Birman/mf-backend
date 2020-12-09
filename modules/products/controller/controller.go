package products

import (
	"net/http"

	"github.com/rvalessandro/mf-backend/modules/products/domain"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/products/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Find(c *gin.Context) {
	products, err := services.FindProduct()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetByCategory(c *gin.Context) {
	categoryID, idErr := parser.ParseID(c.Param("category_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, err := services.GetProductsByCategory(categoryID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Get(c *gin.Context) {
	productID, idErr := parser.ParseID(c.Param("product_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	product, err := services.GetProduct(productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Create(c *gin.Context) {
	var productParam domain.CreateProductParams
	err := c.ShouldBindJSON(&productParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	newProduct, createErr := services.CreateProduct(productParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func Update(c *gin.Context) {
	productID, err := parser.ParseID(c.Param("product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var productParam domain.CreateProductParams
	bindErr := c.ShouldBindJSON(&productParam)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	product, updateErr := services.UpdateProduct(productID, productParam)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Delete(c *gin.Context) {
	productID, err := parser.ParseID(c.Param("product_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var product *domain.Product
	product, err = services.DeleteProduct(productID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, product)
}
