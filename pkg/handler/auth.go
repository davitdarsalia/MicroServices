package handler

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

// @Summary SignUp Handler
// @Tags Auth
// @Description Create An User
// @ID create-user
// @Accept json
// @Produce json
// @Param input body entities.UserRegInput true "User Info"
// @Success 201 {object} entities.RegisteredUserResponse
// @Failure 409 {object} localError
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/sign-up [post]
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

	c.SecureJSON(http.StatusCreated, entities.RegisteredUserResponse{
		UserId:    id,
		Message:   constants.CreatedUserSuccess,
		CreatedAt: time.Now().Format(entities.RegularFormat),
	})
}

// @Summary SignIn Handler
// @Tags Auth
// @Description Login An User
// @ID login-user
// @Accept json
// @Produce json
// @Param input body entities.UserInputWithoutID true "Credentials"
// @Success 200 {object} entities.SignedInUserResponse
// @Failure 404 {object} localError
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/sign-in [post]
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

	c.SecureJSON(http.StatusOK, entities.SignedInUserResponse{
		Message:         constants.SuccessfulSignIn,
		AccessToken:     token,
		AccessTokenExp:  os.Getenv("ACCESS_TOKEN_EXP"),
		RefreshToken:    newRefreshToken(),
		RefreshTokenExp: "Never",
	})

}

// @Summary Reset Password Handler
// @Tags Auth
// @Description Reset Password Outside From Account
// @ID reset-password
// @Accept json
// @Produce json
// @Param input body entities.ResetPassword true "Credentials"
// @Success 205 {object} entities.ResetPasswordResponse
// @Failure 406 {object} localError
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/reset-password [post]
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

	c.JSON(http.StatusResetContent, entities.ResetPasswordResponse{
		UserID:  userId,
		Message: constants.ResetPasswordSuccess,
	})
}

// @Summary Validate Reset Password Mail Handler (Code)
// @Tags Auth
// @Description Validation Of Reset Password (Code Is Sent To Gmail)
// @ID validate-reset-password
// @Accept json
// @Produce json
// @Param input body entities.ValidateResetEmail true "Credentials"
// @Success 205 {object} entities.ValidateResetPasswordResponse
// @Failure 406 {object} localError
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/verify-reset-email [post]
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

// @Summary Reset Password Profile Handler
// @Security ApiKeyAuth
// @Tags Auth
// @Description Reset Password From The Account
// @ID reset-password-profile
// @Accept json
// @Produce json
// @Param input body entities.ValidateResetEmail true "Credentials"
// @Success 205 {object} entities.ValidateResetPasswordResponse
// @Failure 406 {object} localError
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/reset-password-profile [post]n.Context)
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

// @Summary Refresh Login Handler
// @Security ApiKeyAuth
// @Tags Auth
// @Description Refresh Login - Refreshes AccessToken
// @ID refresh-login
// @Accept json
// @Produce json
// @Param input body entities.SignedInUserResponse true "Credentials"
// @Success 200 {object} entities.ValidateResetPasswordResponse
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/refresh-login [post]
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
