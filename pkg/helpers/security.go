package helpers

import (
	"crypto/sha256"
	"fmt"
)

func HashData(d string) string {
	hash := sha256.Sum256([]byte(d))
	str := fmt.Sprintf("%x", hash)
	return str
}
