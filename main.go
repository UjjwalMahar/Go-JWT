package main

import (
	"github.com/UjjwalMahar/go-jwt/controllers"
	"github.com/UjjwalMahar/go-jwt/initializers"
	"github.com/UjjwalMahar/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
	
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth , controllers.Validate)
	r.GET("/health", controllers.Health)
	r.Run()
}