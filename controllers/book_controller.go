package controllers

import (
	"net/http"
	"simple-rest/configs"
	"simple-rest/helpers"
	"simple-rest/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func GetBooks(c *gin.Context) {
	var books models.Books
	tx := configs.DB.Where("title LIKE ?", "%"+c.Query("title")+"%")

	// ?status=1
	if c.Query("status") != "" {
		tx.Where("status = ?", c.Query("status"))
	}

	result := tx.Preload("User").Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Books retrieved",
		"data":    books,
	})
}

type CreateBookRules struct {
	Title       string `form:"title" binding:"required,min=3,max=191"`
	Description string `form:"description" binding:"required,min=3,max=191"`
	Price       int    `form:"price" binding:"required,min=1,max=1000000,numeric"`
}

func StoreBook(c *gin.Context) {
	// Bind form-data with Struct
	var newBook CreateBookRules
	err := c.ShouldBind(&newBook)

	// Show validation errors
	if err != nil {
		errorMessages := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Convert to the base model
	var book models.Book
	copier.Copy(&book, &newBook)

	// Store image if exists
	book.Image = helpers.StoreImage(c, "image")

	book.UserId = 1

	// Save book to db
	result := configs.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created",
		"book":    book,
	})
}

type UpdateBookRules struct {
	Title       string `form:"title" binding:"max=191"`
	Description string `form:"description" binding:"max=191"`
	Price       int    `form:"price" binding:"max=1000000,numeric"`
}

func UpdateBook(c *gin.Context) {
	// Bind form-data with Struct
	var updatedBook UpdateBookRules
	err := c.ShouldBind(&updatedBook)

	// Show validation errors
	if err != nil {
		errorMessages := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Convert to the base model
	var book models.Book
	copier.Copy(&book, &updatedBook)

	// Store image if exists
	book.Image = helpers.StoreImage(c, "image")

	// Save book to db
	result := configs.DB.Where("id = ?", c.Param("book_id")).Updates(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated",
		"book":    book,
	})
}

func DeleteBook(c *gin.Context) {
	result := configs.DB.Delete(&models.Book{}, c.Param("book_id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted",
	})
}
