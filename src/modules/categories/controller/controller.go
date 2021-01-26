package controller

import (
	"github.com/rvalessandro/mf-backend/modules/categories/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rvalessandro/mf-backend/modules/categories/services"
	"github.com/rvalessandro/mf-backend/utils/parser"
)

func Find(c *gin.Context) {
	categories, err := services.FindCategory()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, categories)
}

func Get(c *gin.Context) {
	categoryID, idErr := parser.ParseID(c.Param("category_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	category, err := services.GetCategory(categoryID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func Create(c *gin.Context) {
	categoryParam := domain.NewCategory()
	err := c.ShouldBindJSON(&categoryParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	newCategory, createErr := services.CreateCategory(*categoryParam)
	if createErr != nil {
		c.JSON(http.StatusInternalServerError, createErr)
		return
	}

	c.JSON(http.StatusCreated, newCategory)
}

func Update(c *gin.Context) {
	categoryID, err := parser.ParseID(c.Param("category_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var categoryParam domain.CreateCategoryParams
	bindErr := c.ShouldBindJSON(&categoryParam)
	if bindErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	category, updateErr := services.UpdateCategory(categoryID, categoryParam)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, category)
}

func Delete(c *gin.Context) {
	categoryID, err := parser.ParseID(c.Param("category_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	var category *domain.Category
	category, err = services.DeleteCategory(categoryID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, category)
}
