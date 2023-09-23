package controllers

import (
	"elearning/config"
	"elearning/models"
	"elearning/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = utils.HashAndSalt(user.Password)

	_, err := config.InsertOne("users", user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context)  {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := config.FindOne("users", bson.M{"email": user.Email}, nil)
	if result.Err() != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	pass := utils.ComparePassword(user.Password, []byte(result.Password))


	// Generate an authentication token (we'll implement this later)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}