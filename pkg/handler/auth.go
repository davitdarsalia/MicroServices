package handler

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (h *Handler) signUp(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.RegisterUser(&u)

	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"user_id":    id,
		"message":    "User Created Successfully",
		"created_at": time.Now().Format(entities.RegularFormat),
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Authorization.CheckUser(&u)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id":   id,
		"message":   "User Logged In Successfully",
		"signed_at": time.Now().Format(entities.RegularFormat),
	})
}
