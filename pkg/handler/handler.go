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
			account.GET("/profile-details", h.GetProfileDetails)
			account.GET("/user-info", h.GetUserInfo)
			account.GET("/trusted-devices-list", h.GetTrustedDevices)
			account.GET("/get-user-by-id/{id}", h.GetUserById)
			account.POST("/block-user", h.BlockUser)
			account.GET("/unblock-user", h.UnblockUser)
			account.GET("/blocked-user-list", h.BlockedUsersList)
			account.GET("/upload-profile-image", h.UploadProfileImage)
			account.GET("/logout", h.LogoutSession)
			account.GET("/update-profile-details", h.UpdateProfileDetails)
			account.GET("/update-trusted-devices", h.UpdateTrustedDevices)
			account.GET("/set-passcode", h.SetPasscode)
		}

		settings := protected.Group("/settings", h.checkAuth)
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
