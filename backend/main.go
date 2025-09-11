package main

import (
	"log"
	"mentalhealthchat/db"
	"mentalhealthchat/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Optional: check they are loaded
	dbURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	if dbURL == "" || jwtSecret == "" {
		log.Fatal("DATABASE_URL or JWT_SECRET not set")
	}

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
