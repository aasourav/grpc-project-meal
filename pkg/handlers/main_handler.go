package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MainHandler struct for handling general application routes.
type MainHandler struct{}

// NewMainHandler creates a new instance of MainHandler.
func NewMainHandler() *MainHandler {
	return &MainHandler{}
}

// HealthCheck responds with a simple health status.
func (h *MainHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *MainHandler) AboutUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"name": "meal management", "version": "0.0.1", "developer": "ahsan amin", "email": "ahsan.sourav109@gmail.com"})
}
