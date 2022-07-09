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
}
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
}
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
}
func (h *Handler) updatePrivacySettings(c *gin.Context) {

}
