package service

import (
	"testing"
)

/* Benchmarks */
func BenchmarkAccessToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		accessToken("63e0112d-120f-4d15-be48-a8539ea0218b", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkRefreshToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		refreshToken("63e0112d-120f-4d15-be48-a8539ea0218b", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash("RandomPass", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkGetIPv6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getIPv6()
	}
}

func BenchmarkFormattedDateTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getFormattedDateTime()
	}
}

func BenchmarkGenerateSalt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateSalt()
	}
}

func BenchmarkCheckUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkUUID("123e4567-e89b-12d3-a456-426614174000.")
	}
}
