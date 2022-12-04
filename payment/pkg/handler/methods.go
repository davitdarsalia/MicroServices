package handler

import (
	"github.com/davitdarsalia/payment/internal/constants"
	"github.com/davitdarsalia/payment/internal/entities"
	"github.com/davitdarsalia/payment/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func (h *Handler) Create(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	id, err := h.services.ProviderService.Create(&u)

	if err != nil {
		utils.Error(c, http.StatusConflict, constants.UserAlreadyRegisteredError)
		return
	}

	t := utils.TokenPair(id)

	c.SecureJSON(http.StatusCreated, entities.RegisteredResponse{
		RegisteredUser: entities.RegisteredUser{
			UserId:    id,
			Message:   constants.CreatedUserSuccess,
			CreatedAt: time.Now().Format(constants.RegularFormat),
		},
		Authenticated: entities.Authenticated{
			AT:    t[0],
			ATExp: os.Getenv("exp"),
			RT:    t[1],
			RTExp: "5 Days",
		},
	})
}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) Refresh(c *gin.Context) {

}

func (h *Handler) Reset(c *gin.Context) {
	var r entities.ResetPasswordInput

	if err := c.BindJSON(&r); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	err := h.services.Reset(&r)

	if err != nil {
		utils.Error(c, http.StatusNotFound, constants.UserNotFoundError)
		return
	}

	c.SecureJSON(http.StatusResetContent,
		entities.ResetPasswordResponse{Message: constants.ResetPasswordSuccess},
	)
}

func (h *Handler) Verify(c *gin.Context) {
}
