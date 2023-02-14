package handler

import (
	"auth/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) DefineRoutes() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/authServer")
	{
		auth.POST("/create-user", h.createUser)
		auth.POST("/login-user", h.loginUser)
		auth.POST("/logout-user", h.logoutUser)
		auth.POST("/recover-password", h.recoverPassword)
		auth.POST("/recover-secret-key", h.recoverSecretKey)

		auth.GET("get-user-info", h.getUserInfo)
	}

	return r
}
