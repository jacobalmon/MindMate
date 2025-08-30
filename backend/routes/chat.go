package routes

import (
	"mentalhealthchat/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ChatHandler(c *gin.Context) {
	var req struct {
		UserID int    `json:"user_id"`
		Text   string `json:"text"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := db.InsertMessage(req.UserID, "user", req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save message"})
		return
	}

	// TODO: add AI Response here.
	db.InsertMessage(req.UserID, "bot", "placeholder for AI response")

	c.JSON(http.StatusOK, gin.H{"message": "message_stored"})
}
