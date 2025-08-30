package main

import (
	"mentalhealthchat/db"
	"mentalhealthchat/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect() // Ensure DB is connected before starting server.

	r := gin.Default()

	r.POST("/auth/signup", routes.SignupHandler)
	r.POST("/auth/login", routes.LoginHandler)
	r.POST("/chat/send", routes.ChatHandler)
	r.GET("/chat/history", routes.ChatHistoryHandler)
	r.POST("/mood/submit", routes.MoodHandler)
	r.GET("/mood/history", routes.MoodHistoryHandler)

	r.Run(":8080")
}
