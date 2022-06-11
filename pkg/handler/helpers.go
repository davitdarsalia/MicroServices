package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thanhpk/randstr"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// generateSessionID - Generates unique session ID. Don't edit
func generateSessionID(bytesAmount int) string {
	var saltBytes []byte

	for i := 0; i < 150; i++ {
		saltBytes = randstr.Bytes(bytesAmount)
	}
	return string(saltBytes)
}

func newRefreshToken() string {
	bytes := make([]byte, 55)

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	if _, err := r.Read(bytes); err != nil {
		log.Fatal(err)
	}

	stringRefreshToken := fmt.Sprintf("%x", bytes)

	return stringRefreshToken
}

func LoginSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(constants.SessionID, generateSessionID(35))
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User Logged In - Login Session Created",
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User Signed Out - Sign out Session Created",
	})
}
