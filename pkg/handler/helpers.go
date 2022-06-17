package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/thanhpk/randstr"
	"log"
	"math/rand"
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
