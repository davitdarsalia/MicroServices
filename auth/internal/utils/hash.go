package utils

import (
	"crypto/sha512"
	"fmt"
	"github.com/davitdarsalia/auth/internal/types"
	"github.com/thanhpk/randstr"
	"strings"
)

func Hash(p types.Password, salt string) types.Hash512 {
	hash := sha512.New()
	hash.Write([]byte(p))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func Salt() types.SaltVal {
	var s strings.Builder
	s.Write(randstr.Bytes(25))
	return s.String()
}
