package handler

import (
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/utils"
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
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	tPair, err := h.services.ProviderService.Login(&u)

	if err != nil {
		utils.Error(c, http.StatusNotFound, constants.UserNotFoundError)
		return
	}

	c.SecureJSON(http.StatusOK, entities.LoggedInUserResponse{
		Message: constants.LoggedInUserSuccess,
		Authenticated: entities.Authenticated{
			AT:    tPair[0],
			ATExp: os.Getenv("exp"),
			RT:    tPair[1],
			RTExp: "5 Days",
		},
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	var t entities.RefreshLogin

	if err := c.BindJSON(&t); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	tPair, err := h.services.Refresh(t.RT)

	if err != nil {
		utils.Error(c, http.StatusNotAcceptable, constants.InvalidTokenError)
		return
	}

	c.SecureJSON(http.StatusAccepted, entities.LoggedInUserResponse{
		Message: constants.LoggedInUserSuccess,
		Authenticated: entities.Authenticated{
			AT:    tPair[0],
			ATExp: os.Getenv("exp"),
			RT:    tPair[1],
			RTExp: "5 Days",
		},
	})
}

func (h *Handler) Verify(c *gin.Context) {
	// TODO implement me
}

func (h *Handler) Reset(c *gin.Context) {
	// TODO implement me
}
