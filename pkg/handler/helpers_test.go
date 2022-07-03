package handler

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRefreshToken(t *testing.T) {
	tokenLength := 110

	assert := assert.New(t)

	res := newRefreshToken()

	for i := 0; i < 150; i++ {
		t.Logf("Result: %v ", res)

		if len(res) != tokenLength {
			t.Errorf("Incorrect: Expected Length %v, But Got: %v", tokenLength, "Incorrect Length")
			t.Fail()
		}

		if n := assert.NotNil(res); !n {
			t.Errorf("Incorrect: Expected %v, But Got: %v", "Refresh Token", "Empty String")
			t.Fail()
		}
	}

}

func TestSessionIDManager(t *testing.T) {
	assert := assert.New(t)

	var testCases = []struct {
		input          int
		expectedLength int
	}{
		{
			input:          20,
			expectedLength: 20,
		},
		{
			input:          150,
			expectedLength: 150,
		},
		{
			input:          100,
			expectedLength: 100,
		},
		{
			input:          200,
			expectedLength: 200,
		},
		{
			input:          180,
			expectedLength: 180,
		},
	}

	for _, testCase := range testCases {
		res := generateSessionID(testCase.input)
		t.Logf("Result: %v ", res)

		if len(res) != testCase.expectedLength {
			t.Errorf("Incorrect: Expected %v, But Got: %v", testCase.expectedLength, res)
			t.Fail()
		}

		if n := assert.NotNil(res); !n {
			t.Errorf("Incorrect: Expected %v, But Got: %v", "Session ID", "Empty String")
		}

	}

}
