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
		go auth.POST("/sign-up", h.signUp)
		go auth.POST("/sign-in", h.signIn)
		go auth.POST("/reset-password", h.resetPassword)
		go auth.POST("/verify-reset-email", h.validateResetEmail)

		go auth.POST("/reset-password-profile", h.resetPasswordProfile, h.checkAuth)
		go auth.POST("/refresh-login", h.refreshLogin, h.checkAuth)
	}

	protected := r.Group("api/protected")
	{
		account := protected.Group("/account", h.checkAuth)
		{
			go account.GET("/profile-details", h.GetProfileDetails)
			go account.GET("/user-info", h.GetUserInfo)
			go account.GET("/trusted-devices-list", h.GetTrustedDevices)
			go account.PUT("/add-trusted-device", h.AddTrustedDevice)
			go account.POST("/block-user", h.BlockUser)
			go account.POST("/unblock-user", h.UnblockUser)
			go account.GET("/blocked-user-list", h.BlockedUsersList)
			go account.POST("/upload-profile-image", h.UploadProfileImage)
			go account.POST("/logout", h.LogoutSession)
			go account.PUT("/update-profile-details", h.UpdateProfileDetails)
			go account.POST("/set-passcode", h.SetPasscode)
		}

		settings := protected.Group("/settings", h.checkAuth, h.SessionManager)
		{
			settings.GET("/profile-settings", h.GetProfileSettings)
			settings.GET("/notification-settings", h.GetNotificationSettings)
			settings.GET("/payment-options", h.GetPaymentOptions)
			settings.GET("/privacy-settings", h.GetPrivacySettings)
			settings.GET("/security-settings", h.GetSecuritySettings)
			settings.GET("/update-notification-settings", h.UpdateNotificationSettings)
			settings.GET("/update-payment-options", h.UpdatePaymentOptions)
			settings.GET("/update-privacy-settings", h.UpdatePrivacySettings)
			settings.GET("/update-security-settings", h.UpdateSecuritySettings)
		}

	}

	return r
}
