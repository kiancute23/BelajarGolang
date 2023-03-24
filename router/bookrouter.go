package router

import (
	"golang/controller"

	"github.com/gin-gonic/gin"
)

func StarServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controller.CreateBook)
	router.PUT("/book/:id", controller.UpdateBook)
	router.GET("/book/:id", controller.GetBook)
	router.DELETE("/book:id", controller.DeleteBook)
	return router
}
