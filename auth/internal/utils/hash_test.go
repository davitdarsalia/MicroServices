package utils

import (
	"math/rand"
	"testing"
)

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash(randomString(b), randomString(b))
	}
}

func BenchmarkSalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Salt()
	}
}

// randomString - The helper function, which generates random string for Hashing Benchmark
func randomString(b *testing.B) string {
	b.Helper()
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-!@#$%^&*()_"
	c := charset[rand.Intn(len(charset))]
	return string(c)

}
