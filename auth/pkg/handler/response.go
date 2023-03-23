package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type respError struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)

	//strHeader := (*reflect.StringHeader)(unsafe.Pointer(&message))
	//bytes := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
	//	Data: strHeader.Data,
	//	Len:  strHeader.Len,
	//	Cap:  strHeader.Len,
	//}))
	//
	//if 'a' <= bytes[0] && bytes[0] <= 'z' {
	//	bytes[0] -= 'a' - 'A'
	//}

	c.AbortWithStatusJSON(statusCode, respError{Message: message})
}
