package handler

import (
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) Create(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
	}

	id, err := h.services.ProviderService.Create(u)

	if err != nil {
		utils.Error(c, http.StatusConflict, constants.UserAlreadyRegistered)
		return
	}

	c.SecureJSON(http.StatusCreated, entities.RegisteredUserResponse{
		UserId:    id,
		Message:   constants.CreatedUserSuccess,
		CreatedAt: time.Now().Format(constants.RegularFormat),
	})
}

func (h *Handler) Login(c *gin.Context) {
	// TODO implement me
}

func (h *Handler) Refresh(c *gin.Context) {
	// TODO implement me
}

func (h *Handler) Verify(c *gin.Context) {
	// TODO implement me
}

func (h *Handler) Reset(c *gin.Context) {
	// TODO implement me
}
