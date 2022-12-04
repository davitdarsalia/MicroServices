package handler

import (
	"github.com/davitdarsalia/payment/pkg/service"
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
		auth.POST("/create-user", h.Create)
		auth.POST("/login", h.Login)
		auth.POST("/refresh-login", h.Refresh)

		auth.POST("/reset-password", h.Reset)

		auth.POST("/verify-reset-email", h.Verify)
	}

	return r
}

// New - Returns New And Only Instance Of Handler
func New(s *service.Service) *Handler {
	return &Handler{services: s}
}
