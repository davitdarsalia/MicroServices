package outerServices

import (
	"testing"
)

func BenchmarkAwsCredentials(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetSecret()
	}
}
