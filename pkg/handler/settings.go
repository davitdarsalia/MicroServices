package handler

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getProfileSettings(c *gin.Context) {
}
func (h *Handler) getNotificationSettings(c *gin.Context) {

}
func (h *Handler) getPaymentOptions(c *gin.Context) {

}
func (h *Handler) getPrivacySettings(c *gin.Context) {

}

func (h *Handler) getSecuritySettings(c *gin.Context) {

}

// @Summary Update Notification Settings
// @Tags Protected - Settings
// @Description Updates Notification Settings
// @ID update-notification-settings
// @Accept json
// @Produce json
// @Param input body entities.NotificationSettings true "Credentials"
// @Success 205 {string} constants.UpdateNotificationSettingsSuccess
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/protected/settings/update-notification-settings  [post]
func (h *Handler) updateNotificationSettings(c *gin.Context) {
	var n entities.NotificationSettings

	if err := c.BindJSON(&n); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
	}
	// TODO - Here
	err := h.services.Settings.UpdateNotificationSettings(&n)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.UpdateNotificationSettingsError)
	}

	c.JSON(http.StatusResetContent, constants.UpdateNotificationSettingsSuccess)
}

// @Summary Update Payment Settings
// @Tags Protected - Settings
// @Description Updates Payment Settings
// @ID update-payment-settings
// @Accept json
// @Produce json
// @Param input body entities.PaymentSettings true "Credentials"
// @Success 205 {string} constants.UpdatePaymentSettingsSuccess
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/protected/settings/update-payment-settings  [post]
func (h *Handler) updatePaymentOptions(c *gin.Context) {
	var p entities.PaymentSettings

	if err := c.BindJSON(&p); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
	}
	// TODO - Here
	err := h.services.Settings.UpdatePaymentSettings(&p)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.UpdatePaymentSettingsError)
	}

	c.JSON(http.StatusResetContent, constants.UpdatePaymentSettingsSuccess)
}

// @Summary Update Security Settings
// @Tags Protected - Settings
// @Description Updates Security Settings
// @ID update-security-settings
// @Accept json
// @Produce json
// @Param input body entities.SecuritySettings true "Credentials"
// @Success 205 {string} constants.UpdateSecuritySettingsSuccess
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/protected/settings/update-security-settings  [post]
func (h *Handler) updateSecuritySettings(c *gin.Context) {
	var s entities.SecuritySettings

	if err := c.BindJSON(&s); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
	}
	// TODO - Here
	err := h.services.Settings.UpdateSecuritySettings(&s)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.UpdateSecuritySettingsError)
	}

	c.JSON(http.StatusResetContent, constants.UpdateSecuritySettingsSuccess)
}
func (h *Handler) updatePrivacySettings(c *gin.Context) {

}
