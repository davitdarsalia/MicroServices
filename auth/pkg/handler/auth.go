package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"menuAPI/internal/entities"
	"menuAPI/internal/responses"
	"net/http"
)

func (h *Handler) createUser(c *gin.Context) {
	var u entities.User

	fmt.Println(u)

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.service.CreateUser(u)

	if err != nil {
		newErrorResponse(c, http.StatusConflict, err.Error())
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
}

func (h *Handler) loginUser(c *gin.Context) {
	var u entities.UserInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.service.LoginUser(u)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
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
}
func (h *Handler) recoverPassword(c *gin.Context) {
	var u entities.RecoverPasswordInput

	if err := c.BindJSON(&u); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.RecoverPassword(u)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusResetContent, responses.RecoveredPasswordResponse{
		StatusCode: http.StatusResetContent,
		Message:    responses.RecoveredPasswordSuccessMessage,
	})
}

func (h *Handler) logoutUser(c *gin.Context) {}

func (h *Handler) recoverSecretKey(c *gin.Context) {}

func (h *Handler) getUserInfo(c *gin.Context) {}
