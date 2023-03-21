package service

import (
	"errors"
	"testing"
)

func TestTokenValidator(t *testing.T) {
	t.Parallel()

	t.Run("Expired Tokens", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			token    string
			key      []byte
			expected bool
		}{
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTQ0NjYsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE0MTY2LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6MDk6MjYuNjk4NjEyICswNDAwICswNCBtPSsxOTc4Ljc2NDMzMTQxOCIsInVzZXJfcm9sZSI6IlVzZXIiLCJleHBpcmF0aW9uX3RpbWUiOiI1IE1pbnV0ZXMiLCJpcF9hZGRyZXNzIjoiZmU4MDo6MTgyYTpjNWNhOjNlNjg6Y2Y3MCAifQ.4wMlmjtr1fRpUcVpFSi5sof25wqwuSGevjv43oEdoEI`,
				key:      []byte("tXFJowR5Flf*AP6<oaK*.g7AzTw(:L"),
				expected: false,
			},
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTU3NzksImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE1NzE5LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6MzU6MTkuODk2NTM1ICswNDAwICswNCBtPSszMS45NDY4MTY1MDEiLCJ1c2VyX3JvbGUiOiJVc2VyIiwiZXhwaXJhdGlvbl90aW1lIjoiMSBNaW51dGVzIiwiaXBfYWRkcmVzcyI6ImZlODA6OjE4MmE6YzVjYTozZTY4OmNmNzAgIn0.TPd97agcr-RyWwLapvAnNuILuuerx0x3ISirLSLvqGk`,
				key:      []byte("tXFJowR5Flf*AP6<oaK*.g7AzTw(:L"),
				expected: false,
			},
		}

		for _, c := range cases {
			res, _ := validateToken(c.token, c.key)

			if res != c.expected {
				t.Errorf("Got: %t, Wanted: %t", res, c.expected)
			}
		}
	})

	t.Run("Valid Tokens With Invalid Signature", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			token    string
			key      []byte
			expected bool
		}{
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTU5OTcsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE1OTM3LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6Mzg6NTcuMTcxNzYgKzA0MDAgKzA0IG09KzI0OS4yMzQ2MjUyOTMiLCJ1c2VyX3JvbGUiOiJVc2VyIiwiZXhwaXJhdGlvbl90aW1lIjoiMSBNaW51dGVzIiwiaXBfYWRkcmVzcyI6ImZlODA6OjE4MmE6YzVjYTozZTY4OmNmNzAgIn0.RUDhbjBRO-NhQA9laJFgcnmE_zO9DHdpJSq-KS3m0VI`,
				key:      []byte("df29frj1gj*AP6<oaK*.g7AzTw(:L"),
				expected: false,
			},
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTYwMTAsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE1OTUwLCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6Mzk6MTAuNjYyMTE1ICswNDAwICswNCBtPSsyNjIuNzI1NDcxNjY4IiwidXNlcl9yb2xlIjoiVXNlciIsImV4cGlyYXRpb25fdGltZSI6IjEgTWludXRlcyIsImlwX2FkZHJlc3MiOiJmZTgwOjoxODJhOmM1Y2E6M2U2ODpjZjcwICJ9.QtJ6WM-adek3QUYyrx0uncVDkGLIIV3-JMIq2VS_-ds`,
				key:      []byte("tXFJoww(:L"),
				expected: false,
			},
		}

		for _, c := range cases {
			res, _ := validateToken(c.token, c.key)

			if res != c.expected {
				t.Errorf("Got: %t, Wanted: %t", res, c.expected)
			}
		}
	})

	t.Run("Malformed Inputs", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			token    string
			key      []byte
			expected bool
		}{
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTQ0NjYsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE0MTY2LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6MDk6MjYuNjk4NjEyICswNDAwICswNCBtPSsxOTc4Ljc2NDMzMTQxOCIsInVzZXJfcm9sZSI6IlVzZXIiLCJleHBpcmF0aW9uX3RpbWUiOiI1IE1pbnV0ZXMiLCJpcF9hZGRyZXNzIjoiZmU4MDo6MTgyYTpjNWNhOjNlNjg6Y2Y3MCAifQ.4wMlmjtr1fRpUcVpFSi5sof25wqwuSGevjv43oEdoEI`,
				key:      []byte("tXFJowAzTw(:L"),
				expected: false,
			},
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTU3NzksImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE1NzE5LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6MzU6MTkuODk2NTM1ICswNDAwICswNCBtPSszMS45NDY4MTY1MDEiLCJ1c2VyX3JvbGUiOiJVc2VyIiwiZXhwaXJhdGlvbl90aW1lIjoiMSBNaW51dGVzIiwiaXBfYWRkcmVzcyI6ImZlODA6OjE4MmE6YzVjYTozZTY4OmNmNzAgIn0.TPd97agcr-RyWwLapvAnNuILuuerx0x3ISirLSLvqGk`,
				key:      []byte("tXFJowR5FlTw(:L"),
				expected: false,
			},
		}

		for _, c := range cases {
			res, _ := validateToken(c.token, c.key)

			if res != c.expected {
				t.Errorf("Got: %t, Wanted: %t", res, c.expected)
			}
		}
	})

	t.Run("Correct Inputs", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			token    string
			key      []byte
			expected bool
		}{
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzNDY1MjEsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE2NTIxLCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6NDg6NDEuNzYyODczICswNDAwICswNCBtPSs0LjU5NzI2MDcwOSIsInVzZXJfcm9sZSI6IlVzZXIiLCJleHBpcmF0aW9uX3RpbWUiOiI1MDAgTWludXRlcyIsImlwX2FkZHJlc3MiOiJmZTgwOjoxODJhOmM1Y2E6M2U2ODpjZjcwICJ9.37-C9C0ExEC2P1Q7lKFjADhVgJs8sMfTrjHmgtt8Vzo`,
				key:      []byte("tXFJowR5Flf*AP6<oaK*.g7AzTw(:L"),
				expected: true,
			},
			{
				token:    `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzNDY1MzQsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE2NTM0LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6NDg6NTQuNzk5NTY0ICswNDAwICswNCBtPSsxNy42MzQ0MjY2NjciLCJ1c2VyX3JvbGUiOiJVc2VyIiwiZXhwaXJhdGlvbl90aW1lIjoiNTAwIE1pbnV0ZXMiLCJpcF9hZGRyZXNzIjoiZmU4MDo6MTgyYTpjNWNhOjNlNjg6Y2Y3MCAifQ.1lDC86JpHzxzUfc00CU2wbBKXqTH-x9rjfzJuOji4kA`,
				key:      []byte("tXFJowR5Flf*AP6<oaK*.g7AzTw(:L"),
				expected: true,
			},
		}

		for _, c := range cases {
			res, err := validateToken(c.token, c.key)

			if res != c.expected {
				t.Errorf("Got: %t, Wanted: %t, Error: %s", res, c.expected, err.Error())
			}
		}
	})
}

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

func BenchmarkRefreshToken(b *testing.B) {
	token := `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJSZWd1bGFyIFVzZXIiLCJleHAiOjE2NzkzMTQ0NjYsImlzcyI6IkF1dGggU2VydmVyIiwibmJmIjoxNjc5MzE0MTY2LCJzdWIiOiJBdXRob3JpemF0aW9uLCBBdXRoZW50aWNhdGlvbiIsInRlbF9udW1iZXIiOiIiLCJpZF9udW1iZXIiOiIiLCJ1c2VyX2lkIjoiMmE1YzJhNjEtYWExNS00MzgxLWI3YjUtNGQ2YTdhMjJkMTE4IiwiY3JlYXRlZF9hdCI6IjIwMjMtMDMtMjAgMTY6MDk6MjYuNjk4NjEyICswNDAwICswNCBtPSsxOTc4Ljc2NDMzMTQxOCIsInVzZXJfcm9sZSI6IlVzZXIiLCJleHBpcmF0aW9uX3RpbWUiOiI1IE1pbnV0ZXMiLCJpcF9hZGRyZXNzIjoiZmU4MDo6MTgyYTpjNWNhOjNlNjg6Y2Y3MCAifQ.4wMlmjtr1fRpUcVpFSi5sof25wqwuSGevjv43oEdoEI`
	for i := 0; i < b.N; i++ {
		refreshToken([]byte("tXFJowR5Flf*AP6<oaK*.g7AzTw(:L"), token)
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
