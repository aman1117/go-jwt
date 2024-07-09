package main

import (
	"github.com/aman1117/go-jwt/controllers"
	"github.com/aman1117/go-jwt/initializers"
	"github.com/aman1117/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	r := gin.Default()

	r.GET("/hello", controllers.Hello)
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
