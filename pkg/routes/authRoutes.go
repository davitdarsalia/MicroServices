package routes

import (
	"github.com/davitdarsalia/RestAPI.git/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(e *gin.Engine) {
	auth := e.Group("/api/auth")
	{
		auth.POST("/registration", handlers.RegistrationHandler)
		auth.POST("/login", handlers.LoginHandler)
		auth.POST("/refreshlogin", handlers.RefreshLoginHandler)
		auth.POST("/logout", handlers.LogoutHandler)
	}
}
