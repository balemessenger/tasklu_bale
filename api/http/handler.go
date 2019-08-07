package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"taskulu/pkg"
)

type Handler struct {
	log *pkg.Logger
}

func NewHandler(log *pkg.Logger) *Handler {
	return &Handler{log}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": 1})
	return
}

func (h *Handler) AdminExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello Admin"})
	return
}

func (h *Handler) Example(c *gin.Context) {
	var example Example
	err := c.ShouldBindJSON(&example)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"title": example.Title, "body": example.Body})
	return
}
