package handler

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

	token, err := h.services.Authorization.CheckUser(u.UserName, u.Password)

	// Fix this - err is always nil
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, entities.SignedInUserResponse{
		Message:         "User Successfully Signed In",
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
