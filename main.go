package main

import (
	"github.com/aman1117/go-jwt/controllers"
	"github.com/aman1117/go-jwt/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	// initializers.SyncDatabase()
}
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", controllers.SignUp)
	r.Run()
}