package handler

import (
	"github.com/davitdarsalia/auth/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) Routes() *gin.Engine {
	r := gin.Default()

	docs := r.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := r.Group("api/auth")
	{
		auth.POST("/create-user", h.create)
		auth.POST("/login", h.login)
		auth.POST("/refresh-login", h.refresh)

		auth.POST("/reset-password", h.reset)

		auth.POST("/verify-reset-email", h.verify)
	}

	return r
}

// New - Returns New And Only Instance Of Handler
func New(s *service.Service) *Handler {
	return &Handler{services: s}
}
