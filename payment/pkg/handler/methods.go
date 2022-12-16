package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchPublicKey(c *gin.Context) {

}

type EmailInput struct {
	Email string `json:"email"`
}

type SessionOutput struct {
	Id string `json:"id"`
}

func Checkout() {

}
