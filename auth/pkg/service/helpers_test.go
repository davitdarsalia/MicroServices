package service

import "testing"

func BenchmarkAccessToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = accessToken("63e0112d-120f-4d15-be48-a8539ea0218b", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkRefreshToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = refreshToken("63e0112d-120f-4d15-be48-a8539ea0218b", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hash("RandomPass", "63e0112d-120f-4d15-be48-a8539ea0218b")
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
		_, _ = generateSalt()
	}
}
