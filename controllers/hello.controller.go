package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	fmt.Println("Hello, World!")
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}
