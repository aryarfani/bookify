package controllers

import (
	"bookify/configs"
	"bookify/helpers"
	"bookify/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	var users models.Users
	tx := configs.DB.Where("name LIKE ?", "%"+c.Query("name")+"%")

	// ?status=1
	if c.Query("status") != "" {
		tx.Where("status = ?", c.Query("status"))
	}

	result := tx.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users Retrieved",
		"data":    users,
	})
}

type CreateUserRules struct {
	Name     string `form:"name" binding:"required,max=191"`
	Email    string `form:"email" binding:"required,max=191,email"`
	Password string `form:"password" binding:"required,max=191"`
}

func StoreUser(c *gin.Context) {
	// Bind form-data with Struct
	var newUser CreateUserRules
	err := c.ShouldBind(&newUser)

	// Show validation errors
	if err != nil {
		errorMessages := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Convert to the base model
	var user models.User
	copier.Copy(&user, &newUser)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		user.Password = string(hashedPassword)
	}

	// Save user to db
	result := configs.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
		"user":    user,
	})
}

type UpdateUserRules struct {
	Name     string `form:"name" binding:"max=191"`
	Email    string `form:"email" binding:"omitempty,max=191,email"`
	Password string `form:"password" binding:"max=191"`
	Image    string `form:"image" binding:"omitempty,file"`
}

func UpdateUser(c *gin.Context) {
	// Bind form-data with Struct
	var updatedUser UpdateUserRules
	err := c.ShouldBind(&updatedUser)

	// Show validation errors
	if err != nil {
		errorMessages := helpers.GenerateValidationResponse(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	// Convert to the base model
	var user models.User
	copier.Copy(&user, &updatedUser)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Save user to db
	result := configs.DB.Where("id = ?", c.Param("user_id")).Updates(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	result := configs.DB.Delete(&models.User{}, c.Param("user_id"))
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted",
	})
}
