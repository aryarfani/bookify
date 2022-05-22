package main

import (
	"bookify/configs"
	"bookify/controllers"
	"bookify/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB init
	configs.ConnectDB()
	configs.DB.AutoMigrate(&models.Book{}, &models.User{})

	// Router init
	router := gin.Default()
	router.Static("/images", "./images")

	v1 := router.Group("/v1")

	v1.GET("/books", controllers.GetBooks)
	v1.POST("/books", controllers.StoreBook)
	v1.PUT("/books/:book_id", controllers.UpdateBook)
	v1.DELETE("/books/:book_id", controllers.DeleteBook)

	v1.GET("/users", controllers.GetUsers)
	v1.POST("/users", controllers.StoreUser)
	v1.PUT("/users/:user_id", controllers.UpdateUser)
	v1.DELETE("/users/:user_id", controllers.DeleteUser)

	router.Run()
}
