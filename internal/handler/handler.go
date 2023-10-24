package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/tonytangdev/go-rest-api/internal/model"
	"gorm.io/gorm"
)

type UserInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type UserUpdateInput struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

var validate = validator.New()

func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB) // Retrieve and assert db from context
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	if result := db.Create(&newUser); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func DeleteUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get user ID from URL
	userID := c.Param("id")

	// Find and delete the user
	var user model.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	db.Delete(&user)

	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get user ID from URL
	userID := c.Param("id")

	// Validate request body
	var input UserUpdateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find and update the user
	var user model.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	db.Save(&user)

	c.JSON(http.StatusOK, user)
}
