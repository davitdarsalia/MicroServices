package service

import (
	"auth/internal/entities"
	"errors"
	"testing"
)

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash("RandomPassword", "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkGenerateStruct(b *testing.B) {
	input := errors.New("verifications failed for fields: [Email]")

	for i := 0; i < b.N; i++ {
		generateValidationStruct(input)
	}
}

func BenchmarkAccessToken(b *testing.B) {
	user := entities.User{
		Name:      "Random",
		Surname:   "User",
		UserName:  "Random.User",
		Email:     "random21@gmail.com",
		TelNumber: "+995599223215",
		IDNumber:  "21212121521",
		Password:  "RandomUserPass",
		Salt:      "63e0112d",
	}
	for i := 0; i < b.N; i++ {
		accessToken([]byte("63e0112d-120f-4d15-be48-a8539ea0218b"), &user, "63e0112d-120f-4d15-be48-a8539ea0218b")
	}
}

func BenchmarkRefreshToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		refreshToken()
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
		salt()
	}
}

func BenchmarkCheckUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkUUID("123e4567-e89b-12d3-a456-426614174000.")
	}
}
