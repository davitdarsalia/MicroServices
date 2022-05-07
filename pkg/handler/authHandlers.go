package handler

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var userInstance entities.User

	if err := c.BindJSON(&userInstance); err != nil {
		logrus.Fatalf("Error")
	}
}

func (h *Handler) signIn(c *gin.Context) {

}
