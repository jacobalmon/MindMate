package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Temp signup handler.
func SignupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "signup endpoint working"})
}

// Temp login handler.
func LoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login endpoint working"})
}
