package routes

import (
	"mentalhealthchat/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MoodHandler(c *gin.Context) {
	var req struct {
		UserID int    `json:"user_id"`
		Note   string `json:"note"`
		Score  int    `json:"mood_score"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Note == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "note cannot be empty"})
		return
	}

	if req.Score < 1 || req.Score > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "score must be between 1 and 5"})
		return
	}

	err := db.InsertMood(req.UserID, req.Note, req.Score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save mood"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mood stored"})
}
