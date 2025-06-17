package handler

import (
	"github.com/TimmyTurner98/sharing/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var input models.UserRegister

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	id, err := h.services.Auth.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
