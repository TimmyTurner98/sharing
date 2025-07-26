package handler

import (
	"github.com/TimmyTurner98/sharing/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	{
		apiV1 := router.Group("api/v1")
		apiV1.POST("/test", testhandler)
		apiV1.POST("/login", h.SendCode)
		apiV1.POST("/verify", h.VerifyCode)
	}

	return router
}
