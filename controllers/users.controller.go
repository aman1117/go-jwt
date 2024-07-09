package controllers

import (
	"fmt"
	"net/http"

	"github.com/aman1117/go-jwt/initializers"
	"github.com/aman1117/go-jwt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	fmt.Printf("body: %v", body)

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Parameters"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while hashing password"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error while creating user"})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}
