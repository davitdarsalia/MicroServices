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

	c.JSON(http.StatusCreated, entities.RegisteredUserResponse{
		UserId:    id,
		Message:   "User Created Successfully",
		CreatedAt: time.Now().Format(entities.RegularFormat),
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusNotAcceptable, err.Error())
		return
	}

	userID, err := h.services.CheckUser(&u)

	// userID - If 0 , means that user doesn't exists
	if userID == 0 {
		c.JSON(http.StatusNotFound, entities.SignedInUserResponse{
			Message: "User Not Found",
		})
		return
	}
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	token, tokenError := generateToken(userID)

	if tokenError != nil {
		return
	}

	c.JSON(http.StatusOK, entities.SignedInUserResponse{
		UserId:      userID,
		Message:     "User Signed In Successfully",
		AccessToken: token,
	})
}
