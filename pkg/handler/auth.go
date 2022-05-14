package handler

import (
	"fmt"
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

	c.JSON(http.StatusCreated, entities.RegisteredUserResponse{
		UserId:    id,
		Message:   "User Created Successfully",
		CreatedAt: time.Now().Format(entities.RegularFormat),
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.CheckUser(&u)
	fmt.Println(user)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{})
}
