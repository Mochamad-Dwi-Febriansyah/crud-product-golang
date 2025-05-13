package main

import (
	"gin/controllers/authcontroller"
	"gin/controllers/productcontrollers"
	"gin/database"
	"gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.ConnectToDB()

	// Create a Gin router
	r := gin.Default()

	r.Use(func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Auth routes
	r.POST("/login", authcontroller.Login)
	r.POST("/register", authcontroller.Register)
	r.GET("/logout", authcontroller.Logout)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware()) // JWT middleware diaktifkan untuk semua route dalam /api
	{
		api.GET("/products", productcontrollers.GetAll)
		api.GET("/products/:id", productcontrollers.GetByID)
		api.POST("/products", productcontrollers.Create)
		api.PUT("/products/:id", productcontrollers.Update)
		api.DELETE("/products/:id", productcontrollers.Delete)
	}

	r.Static("/uploads", "./uploads")
	// Start the server
	r.Run(":8080")
}
