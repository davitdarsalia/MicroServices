package handler

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func (h *Handler) signUp(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	id, err := h.services.Authorization.RegisterUser(&u)

	if err != nil {
		newErrorResponse(c, http.StatusConflict, constants.UserAlreadyRegistered)
		return
	}

	c.JSON(http.StatusCreated, entities.RegisteredUserResponse{
		UserId:    id,
		Message:   constants.CreatedUserSuccess,
		CreatedAt: time.Now().Format(entities.RegularFormat),
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	token, err := h.services.Authorization.CheckUser(u.UserName, u.Password)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, constants.UserNotFoundError)
		return
	}

	c.JSON(http.StatusOK, entities.SignedInUserResponse{
		Message:         constants.SuccessfulSignIn,
		AccessToken:     token,
		AccessTokenExp:  os.Getenv("ACCESS_TOKEN_EXP"),
		RefreshToken:    newRefreshToken(),
		RefreshTokenExp: "Never",
	})

}

func (h *Handler) refreshLogin(c *gin.Context) {
}
func (h *Handler) resetPassword(c *gin.Context) {
}
func (h *Handler) resetPasswordProfile(c *gin.Context) {
}
func (h *Handler) otpGenerator(c *gin.Context) {
}
