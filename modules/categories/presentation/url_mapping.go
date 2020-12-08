package presentation

import (
	"github.com/gin-gonic/gin"
	categoryController "github.com/rvalessandro/mf-backend/modules/categories/controller"
)

func MapURLs(router *gin.Engine) {
	router.GET("/categories", categoryController.Find)
	router.GET("/categories/:category_id", categoryController.Get)
	router.POST("/categories", categoryController.Create)
	router.PUT("/categories/:category_id", categoryController.Update)
	router.DELETE("/categories/:category_id", categoryController.Delete)
}
