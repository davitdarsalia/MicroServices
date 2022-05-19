package handler

import (
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	auth := r.Group("/api/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/reset-password", h.resetPassword)
		auth.POST("/verify-reset-email", h.validateResetEmail)

		auth.POST("/reset-password-profile", h.resetPasswordProfile, h.checkAuth)
		auth.POST("/refresh-login", h.refreshLogin, h.checkAuth)
	}

	protected := r.Group("api/protected")
	{
		account := protected.Group("/account", h.checkAuth)
		{
			account.GET("/user-info", h.signUp)
		}

	}

	return r
}
