package main

import (
	"example/kwame/first_web/controllers"
	"example/kwame/first_web/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	// })

	r.PATCH("/books/:id", controllers.UpdateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
