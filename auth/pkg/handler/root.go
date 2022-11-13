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
		auth.POST("/create-user", h.createUser)
		auth.POST("/login", h.createUser)
		auth.POST("/refresh-login", h.createUser)

		auth.POST("/reset-password", h.createUser)

		auth.POST("/verify-reset-email", h.createUser)
	}

	return r
}

// NewHandler - Returns New And Only Instance Of Handler
func New(s *service.Service) *Handler {
	return &Handler{services: s}
}
