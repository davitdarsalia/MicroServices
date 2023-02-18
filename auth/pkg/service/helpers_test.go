package service

import (
	"strings"
	"testing"
)

// TestTokenChecker - Applies for both, access and refresh tokens
func TestTokenChecker(t *testing.T) {
	t.Run("All inputs are correct", func(t *testing.T) {
		cases := []struct {
			token          string
			key            string
			expectedUserID string
			expected       bool
		}{
			{
				token:          `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc0MzI4OTQsImp0aSI6IjRmYzQyZDliLTU1NjYtNGVhNy05YmVmLTA3Yzk1NTZiMjNjNyIsImlhdCI6MTY3NjcxMjg5NCwic3ViIjoiQXV0aGVudGljYXRpb24gIn0.TB6iJcIvZFNmOaqHFn8LJPg0xeWEP9dYWqUEtzU_jXc`,
				key:            `Cl8jfHhhkKctkNwtTymqFghkK5PNfhx6FWhfAPPt`,
				expectedUserID: `4fc42d9b-5566-4ea7-9bef-07c9556b23c7`,
				expected:       true,
			},
		}

		for _, c := range cases {
			userID, err := checkToken(c.token, c.key)

			if c.expectedUserID != userID || err != nil {
				t.Errorf("Error: %s", err.Error())
			}
		}
	})

	t.Run("Malformed inputs", func(t *testing.T) {
		cases := []struct {
			token          string
			key            string
			expectedUserID string
			expected       bool
		}{
			{
				token:          `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Nzc0MzI4OTQsImp0aSI6IjRmYzQyZDliLTU1NjYtNGVhNy05YmVmLTA3Yzk1NTZiMjNjNyIsImlhdCI6MTY3NjcxMjg5NCwic3ViIjoiQXV0aGVudGljYXRpb24gIn0.TB6iJcIvZFNmOaqHFn8LJPg0xeWEP9dYWqUEtzU_jXc`,
				key:            `Cl8jf1234ctkNwtTy12mqFghkK5P1Nfhx6FWhfAPPt`,
				expectedUserID: `4fc42d9b-5566-4ea7-ds9bef-07c9556b23c7`,
				expected:       false,
			},
		}

		for _, c := range cases {
			userID, err := checkToken(c.token, c.key)

			if c.expectedUserID == userID || err == nil {
				t.Errorf("Impossible scenario. Inputs are maleformed initially")
			}

		}
	})
}

// TestTokens - Applies for both, access and refresh tokens
func TestTokens(t *testing.T) {
	t.Run("Correct Inputs", func(t *testing.T) {
		cases := []struct {
			userID string
			key    string
		}{
			{
				userID: `024cd634-c998-459b-a46e-580da56b4c81`,
				key:    `Xy6s6T7dWRF0-zWhoxP-OyYawfb8dlAW3ST-ldFr`,
			},
			{
				userID: `4fc42d9b-5566-4ea7-9bef-07c9556b23c7`,
				key:    `Cl8jfHhhkKctkNwtTymqFghkK5PNfhx6FWhfAPPt`,
			},
		}

		for _, c := range cases {
			generatedToken, err := accessToken(c.userID, c.key)

			if err != nil || generatedToken == "" {
				t.Errorf("Token generation error: %s", err.Error())
				return
			}

			_, err = checkToken(generatedToken, c.key)

			if err != nil {
				t.Errorf("Token checker error: %s", err.Error())
				return
			}
		}
	})

	t.Run("Malformed inputs", func(t *testing.T) {
		cases := []struct {
			userID string
			key    string
		}{
			/*
				In the first case userID is malformed,
				in the second case, key.
				The Purpose is to test token generator and case must be failed,
				even if one of the inputs is correct format
			*/
			{
				userID: `024cd634-c998s580da56b4c81`,
				key:    `Xy6s6T7dWRF0-zWhoxP-OyYawfb8dlAW3ST-ldFr`,
			},
			{
				userID: `4fc42d9b-5566-4ea7-9bef-07c9556b23c7`,
				key:    `Cl8jhx6FWhfAPPt`,
			},
			{
				userID: `4fc44ea7-9bef-07c92556b23c7`,
				key:    `Cl8jhx6FWhfAPPt`,
			},
		}

		for _, c := range cases {
			generatedToken, err := accessToken(c.userID, c.key)

			if generatedToken != "" {
				t.Error("Token must be empty, because inputs are malformed")
				return
			}

			_, err = checkToken(generatedToken, c.key)

			if err == nil {
				t.Errorf("Impossible scenario. Inputs are maleformed, error must exist")
				return
			}
		}
	})
}

func TestUUIDChecker(t *testing.T) {
	cases := []struct {
		uuid     string
		expected bool
	}{
		{
			uuid:     `8b22b6cf-3e3f-4c7b-a06c-6a447e1d2c98`,
			expected: true,
		}, {
			uuid:     `9aa25c36-144f-467f-a1f7-8bde875c7f3e`,
			expected: true,
		}, {
			uuid:     `7a8cf132-1c7d-4745-b17e-0088f5b5d5a5`,
			expected: true,
		}, {
			uuid:     `5a8469c1-87e7-4d33-bf58-bb8008b2d5a5`,
			expected: true,
		}, {
			uuid:     `2fa15b9f-5e44-4d1b-ae1f-8c87e69d3b67`,
			expected: true,
		}, {
			uuid:     `d95c4c4d-4a8e-4f08-87dd-e854a6ed7f6b`,
			expected: true,
		}, {
			uuid:     `d95c1fc44a8e-4f08-87dd-e854a6ed7f6b`,
			expected: false,
		}, {
			uuid:     `d95c4c4d-4asdafdasf1ed7f6b`,
			expected: false,
		}, {
			uuid:     `d9dsaf2f54e854a6ed7f6b`,
			expected: false,
		},
	}

	for index, c := range cases {
		got := checkUUID(c.uuid)

		if got != c.expected {
			t.Errorf("Test N%d - Got: %v, Wanted: %s", index, got, c.uuid)
		}
	}
}

func TestHash(t *testing.T) {
	cases := []struct {
		text              string
		cryptographicSalt string
	}{
		{
			text:              "randomPassword",
			cryptographicSalt: "jf93g29fwcsjf93g29fwcsjf93g29fwcsjf93g29fwcsjf93g29fwcs",
		},
	}

	for _, c := range cases {
		hash := hash(c.text, c.cryptographicSalt)

		if len(strings.Split(hash, "")) < 170 || len(strings.Split(hash, "")) > 300 {
			t.Error("Incorrect Size Of Hash")
		}
	}
}

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
		_, _ = generateSalt()
	}
}
