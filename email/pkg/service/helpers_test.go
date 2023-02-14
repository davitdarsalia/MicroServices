package service

import (
	"fmt"
	"testing"
)

func TestOTP(t *testing.T) {

	for i := 0; i < 13000; i++ {
		otp := otp()
		strOTPLength := len(fmt.Sprintf("%d", otp))

		if otp < 0 || otp < 100000 || otp > 999999 || strOTPLength < 6 {
			t.Errorf("OTP value is out of range.\nIndex: %d, Value: %d, Length: %d",
				i, otp, strOTPLength,
			)
		}
	}
}

func BenchmarkOTP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		otp()
	}
}
