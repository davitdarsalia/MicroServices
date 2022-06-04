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
func (h *Handler) resetPassword(c *gin.Context) {
	var r entities.ResetPassword

	if err := c.BindJSON(&r); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	userId, err := h.services.Authorization.ResetPassword(&r)

	if err != nil {
		newErrorResponse(c, http.StatusNotAcceptable, constants.ResetPasswordError)
		return
	}

	c.JSON(http.StatusOK, entities.ResetPasswordResponse{
		UserID:  userId,
		Message: constants.ResetPasswordSuccess,
	})
}

func (h *Handler) validateResetEmail(c *gin.Context) {
	var e entities.ValidateResetEmail

	if err := c.BindJSON(&e); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	err := h.services.ValidateResetEmail(&e)

	if err != nil {
		newErrorResponse(c, http.StatusNotAcceptable, constants.ValidateResetPasswordError)
		return
	}

	c.JSON(http.StatusResetContent, entities.ValidateResetPasswordResponse{
		Message:   constants.ValidateResetPasswordSuccess,
		ResetDate: time.Now().Format(entities.RegularFormat),
	})
}

// Start Here
func (h *Handler) resetPasswordProfile(c *gin.Context) {
	var refreshInput entities.ResetPasswordInput

	if err := c.BindJSON(&refreshInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	err := h.services.ResetPasswordProfile(&refreshInput)

	if err != nil {
		newErrorResponse(c, http.StatusNotAcceptable, constants.ResetPasswordError)
		return
	}

	c.JSON(http.StatusResetContent, entities.ResetPasswordProfileResponse{
		Message:   constants.ValidateResetPasswordSuccess,
		ResetDate: time.Now().Format(entities.RegularFormat),
	})
}

func (h *Handler) refreshLogin(c *gin.Context) {
	var r entities.RefreshLoginInput

	if err := c.BindJSON(&r); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	id := h.services.RefreshLogin()

	token, err := entities.GenerateToken(id)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.InternalServerError)
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
func (h *Handler) otpGenerator(c *gin.Context) {
}
