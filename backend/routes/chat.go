package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChatHandler is a test endpoint
func ChatHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "chat endpoint working"})
}
