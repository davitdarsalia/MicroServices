package handler

import "testing"

func BenchmarkMiddleware(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CustomLogger()
	}
}
