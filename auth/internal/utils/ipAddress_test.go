package utils

import "testing"

func BenchmarkIpAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IpAddress()
	}
}
