package routes

import (
	"context"
	"mentalhealthchat/db"
	"net/http"
	"strconv"
	"time"

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

func MoodHistoryHandler(c *gin.Context) {
	type Mood struct {
		ID        int       `json:"id"`
		Note      string    `json:"note"`
		Score     int       `json:"mood_score"`
		CreatedAt time.Time `json:"created_at"`
	}

	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 50 // default last 50 messages
	}

	rows, err := db.Pool.Query(context.Background(),
		"SELECT id, mood_score, note, created_at FROM moods WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2",
		userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch moods"})
		return
	}
	defer rows.Close()

	var moods []Mood
	for rows.Next() {
		var m Mood
		if err := rows.Scan(&m.ID, &m.Score, &m.Note, &m.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not scan mood"})
			return
		}
		moods = append(moods, m)
	}

	for i, j := 0, len(moods)-1; i < j; i, j = i+1, j-1 {
		moods[i], moods[j] = moods[j], moods[i]
	}

	c.JSON(http.StatusOK, gin.H{"moods": moods})
}
