package handler

import (
	"auth/internal/entities"
	"auth/internal/responses"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// createUser creates a new user.
// @Summary Create a new user
// @Description Creates a new user with the given details.
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User details"
// @Success 201 {object} responses.CreateUserResponse
// @Failure 400 {object} newErrorResponse
// @Failure 406 {object} newErrorResponse
// @Failure 409 {object} newErrorResponse
// @Router /create-user [post]
func (h *Handler) createUser(c *gin.Context) {
	var u entities.User

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, responses.BadRequestErrorMessage)
		return
	}

	resp, err := h.service.CreateUser(&u)

	if err != nil {
		var statusCode int

		if strings.Contains(err.Error(), responses.ValidationFailedErrorMessage) {
			statusCode = http.StatusNotAcceptable
		} else {
			statusCode = http.StatusConflict
		}
		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusCreated, responses.CreateUserResponse{
		UserID: resp.UserID,
		CreateUserGenericMessage: responses.CreateUserGenericMessage{
			StatusCode: http.StatusCreated,
			Message:    responses.CreateUserSuccessMessage,
		},
		AuthenticatedUserResponse: entities.AuthenticatedUserResponse{
			AccessToken:           resp.AccessToken,
			AccessTokenExpiresAt:  resp.AccessTokenExpiresAt,
			RefreshToken:          resp.RefreshToken,
			RefreshTokenExpiresAt: resp.RefreshTokenExpiresAt,
		},
	})

	c.Request.Body.Close()
}

// @Summary Log in an existing user
// @Description Logs in an existing user and returns an access and refresh token
// @Tags authentication
// @Accept json
// @Produce json
// @Param user body UserInput true "User login information"
// @Success 200 {object} LoginUserResponse
// @Failure 400 {object} respError
// @Failure 404 {object} respError
// @Router /login-user [post]
func (h *Handler) loginUser(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, responses.BadRequestErrorMessage)
		return
	}

	resp, err := h.service.LoginUser(u)

	if err != nil {
		var statusCode int

		if strings.Contains(err.Error(), responses.ValidationFailedErrorMessage) {
			statusCode = http.StatusNotAcceptable
		} else {
			statusCode = http.StatusNotFound
		}

		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.LoginUserResponse{
		UserID: resp.UserID,
		LoginUserGenericMessage: responses.LoginUserGenericMessage{
			StatusCode: http.StatusOK,
			Message:    responses.LoggedInUserSuccessMessage,
		},
		AuthenticatedUserResponse: entities.AuthenticatedUserResponse{
			AccessToken:           resp.AccessToken,
			AccessTokenExpiresAt:  resp.AccessTokenExpiresAt,
			RefreshToken:          resp.RefreshToken,
			RefreshTokenExpiresAt: resp.RefreshTokenExpiresAt,
		},
	})

	c.Request.Body.Close()
}

func (h *Handler) requestResetPassword(c *gin.Context) {
	var u entities.RecoverPasswordInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, responses.BadRequestErrorMessage)
		return
	}

	err := h.service.RequestPasswordRecover(&u)

	if err != nil {
		var statusCode int

		if strings.Contains(err.Error(), responses.ValidationFailedErrorMessage) {
			statusCode = http.StatusNotAcceptable
		} else {
			statusCode = http.StatusNotFound
		}

		newErrorResponse(c, statusCode, err.Error())
		return
	}

	c.JSON(http.StatusOK, responses.RecoveredPasswordResponse{
		StatusCode: http.StatusOK,
		Message:    responses.RecoveredPasswordSuccessMessage,
	})
	c.Request.Body.Close()
}

func (h *Handler) resetPassword(c *gin.Context) {
	c.Request.Body.Close()
}

func (h *Handler) logoutUser(c *gin.Context) {
	c.Request.Body.Close()
}

func (h *Handler) getUserInfo(c *gin.Context) {
	c.JSON(200, "DDD")

	c.Request.Body.Close()
}
