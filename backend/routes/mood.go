package routes

import (
	"mentalhealthchat/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoodHandler(c *gin.Context) {
	var req struct {
		UserID int    `json:"user_id"`
		Mood   string `json:"mood"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := db.InsertMood(req.UserID, req.Mood)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save mood"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mood stored"})
}
