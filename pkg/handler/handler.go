package handler

import (
	_ "github.com/davitdarsalia/LendAppBackend/docs"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	swaggerDocs := r.Group("/docs")
	{
		swaggerDocs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
			account.GET("/profile-details", h.getProfileDetails)
			account.GET("/user-info", h.getUserInfo)
			account.GET("/trusted-devices-list", h.getTrustedDevices)
			account.GET("/get-images", h.getImages)
			account.PUT("/add-trusted-device", h.addTrustedDevice)
			account.POST("/block-user", h.blockUser)
			account.POST("/unblock-user", h.unblockUser)
			account.GET("/blocked-user-list", h.blockedUsersList)
			account.POST("/upload-profile-image", h.uploadProfileImage)
			account.POST("/logout", h.logoutSession)
			account.PUT("/update-profile-details", h.updateProfileDetails)
			account.POST("/set-passcode", h.setPasscode)
		}

		settings := protected.Group("/settings", h.checkAuth, h.SessionManager)
		{
			settings.GET("/profile-settings", h.getProfileSettings)
			settings.GET("/notification-settings", h.getNotificationSettings)
			settings.GET("/payment-options", h.getPaymentOptions)
			settings.GET("/privacy-settings", h.getPrivacySettings)
			settings.GET("/security-settings", h.getSecuritySettings)
			settings.PUT("/update-notification-settings", h.updateNotificationSettings)
			settings.PUT("/update-payment-options", h.updatePaymentOptions)
			settings.PUT("/update-privacy-settings", h.updatePrivacySettings)
			settings.PUT("/update-security-settings", h.updateSecuritySettings)
		}

	}

	return r
}
