package handler

import (
	_ "auth/cmd/docs"
	"auth/pkg/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:generate mockery --name=Handler
type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) DefineRoutes() *gin.Engine {
	r := gin.New()
	r.Use(customLogger(), sizeLimiter())

	docs := r.Group("/docs")
	{
		docs.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := r.Group("/authServer")
	{
		auth.POST("/create-user", h.createUser)
		auth.POST("/login-user", h.loginUser)
		auth.POST("/logout-user", h.logoutUser)

		auth.POST("/request-password-reset", h.requestResetPassword)
		auth.POST("/reset-password", h.resetPassword)
	}

	return r
}
