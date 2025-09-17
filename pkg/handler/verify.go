package handler

import (
	"net/http"

	"github.com/TimmyTurner98/sharing/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) VerifyCode(c *gin.Context) {
	var input models.VerifyCode

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	accessToken, refreshToken, err := h.services.Auth.VerifyCode(input.Number, input.Code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
