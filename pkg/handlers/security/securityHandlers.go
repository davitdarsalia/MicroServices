package security

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashData - Uses Sha256 as an encryption algorithm.
func HashData(data string) string {
	hash := sha256.Sum256([]byte(data))
	strData := fmt.Sprintf("%x", hash)

	return strData
}

// CompareHashes - Compares hashed values

func CompareHashes(data, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))

	if err != nil {
		log.Println("Error While Comparing Hashes")
	}

	return err == nil
}
