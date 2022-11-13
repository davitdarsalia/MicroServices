package handler

import (
	"github.com/davitdarsalia/auth/internal/constants"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		utils.Error(c, http.StatusBadRequest, constants.BadRequest)
	}

	//id, err := h.services.ProviderService.
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
