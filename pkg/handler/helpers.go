package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"log"
	"math/rand"
	"time"
)

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
