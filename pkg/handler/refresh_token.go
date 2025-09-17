package handler

import (
	"github.com/TimmyTurner98/sharing/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Refresh(c *gin.Context) {
	var input models.RefreshInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	access, refresh, err := h.services.Auth.RefreshTokens(input)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}
