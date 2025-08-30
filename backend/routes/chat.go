package routes

import (
	"context"
	"mentalhealthchat/db"
	"net/http"
	"strconv"
	"time"

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

	err := db.InsertMessage(req.UserID, "users", req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save message"})
		return
	}

	// TODO: add AI Response here.
	db.InsertMessage(req.UserID, "ai", "placeholder for AI response")

	c.JSON(http.StatusOK, gin.H{"message": "message_stored"})
}

func ChatHistoryHandler(c *gin.Context) {
	type Message struct {
		ID        int       `json:"id"`
		Role      string    `json:"role"`
		Text      string    `json:"text"`
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
		"SELECT id, role, text, created_at FROM messages WHERE user_id=$1 ORDER BY created_at DESC LIMIT $2",
		userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch messages"})
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Role, &m.Text, &m.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not scan message"})
			return
		}
		messages = append(messages, m)
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
