package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type localError struct {
	Message string `json:"message"`
}

func Error(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, localError{Message: message})
}
