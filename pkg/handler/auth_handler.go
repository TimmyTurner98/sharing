package handler

import (
	"errors"
	"github.com/TimmyTurner98/sharing/models"
	"github.com/TimmyTurner98/sharing/pkg/service"
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
	if errors.Is(err, service.ErrInvalidNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid phone number format"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
