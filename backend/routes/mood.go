package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MoodHandler is a test endpoint
func MoodHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "mood endpoint working"})
}
