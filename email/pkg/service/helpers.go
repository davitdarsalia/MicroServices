package service

import (
	"crypto/rand"
	"log"
	"math/big"
)

func otp() int64 {
	r, err := rand.Int(rand.Reader, big.NewInt(89999))

	if err != nil {
		log.Printf("Failed to generate otp: %s", err.Error())
	}

	return r.Int64() + 100000
}
