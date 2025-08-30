package main

import (
	"mentalhealthchat/db"
	"mentalhealthchat/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect() // Ensure DB is connected before starting server.

	r := gin.Default()

	// Public Routes.
	r.POST("/auth/signup", routes.SignupHandler)
	r.POST("/auth/login", routes.LoginHandler)

	// Protected Routes.
	auth := r.Group("/")
	auth.Use(routes.AuthMiddleware())
	auth.POST("/chat/send", routes.ChatHandler)
	auth.GET("/chat/history", routes.ChatHistoryHandler)
	auth.POST("/mood/submit", routes.MoodHandler)
	auth.GET("/mood/history", routes.MoodHistoryHandler)

	r.Run(":8080")
}
