package service

import (
	"math/rand"
	"testing"
	"time"
)

var authServiceInstance = NewAuthService(nil, nil)

var randInstance = rand.Intn(150000)

func TestTokenGenerator(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		token, err := authServiceInstance.GenerateToken(randInstance)

		if len(token) < 273 && len(token) > 321 && err != nil {
			t.Errorf("Incorrect Type Of Token")
		} else {
			t.Logf("Input %v, Result Is: %v", randInstance, "Positive")
		}
	}

}

func TestTokenParser(t *testing.T) {
	testCases := []struct {
		sampleToken string
	}{
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjEwNTMyNiIsImlhdCI6MTY1Njg4MTk1OSwic3ViIjoiQXV0aGVudGljYXRpb24iLCJ1c2VyX2lkIjoxMDUzMjYsInVzZXJuYW1lIjoiIiwicm9sZSI6InVzZXIiLCJpcF9hZGRyZXNzIjoiSVB2NDogMTI3LjAuMC4xICJ9.qVIYZYwp5isvhay_F36MikJGGZHivOT0yvARNV1F8vQ",
		},
		{
			sampleToken: " eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjgwODA3IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjgwODA3LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.6hE6ipPiFXQfe4CzcSvfyq0seql8JXwZjj4NFBb3zfk",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjEwNjg4IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjEwNjg4LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.rfsqiSSbWq3Prxbob8E5d0XqC-fDe_FUvclXQD3HrbI",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjIzMzk3IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjIzMzk3LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.kjejBoBK9pfGHsgnZ-tPs1_XU-m5gQjRUcqvpq702Ak",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjQ0Mzc5IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjQ0Mzc5LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.NeS1TrvitMt68bRy7qw5lyfkBFhbrwXzOWDRElr2ELo",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjE0MTM0NiIsImlhdCI6MTY1Njg4MTk1OSwic3ViIjoiQXV0aGVudGljYXRpb24iLCJ1c2VyX2lkIjoxNDEzNDYsInVzZXJuYW1lIjoiIiwicm9sZSI6InVzZXIiLCJpcF9hZGRyZXNzIjoiSVB2NDogMTI3LjAuMC4xICJ9.Z8GZwjizx2YXdnSSWH2zcnb1sZJbqDVo_MvPBVEsF6o",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjU5Nzk0IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjU5Nzk0LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.s6dnoRiEB7QnOrYq-SGB-bCxta6SVX-kS0tWgGvofVk",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6Ijc2Nzk5IiwiaWF0IjoxNjU2ODgxOTU5LCJzdWIiOiJBdXRoZW50aWNhdGlvbiIsInVzZXJfaWQiOjc2Nzk5LCJ1c2VybmFtZSI6IiIsInJvbGUiOiJ1c2VyIiwiaXBfYWRkcmVzcyI6IklQdjQ6IDEyNy4wLjAuMSAifQ.-mA1E1jF_HrZP-CE6xxaak1yaacc7Umj3MBB1UiMaLQ",
		},
		{
			sampleToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTY4ODE5NTksImp0aSI6IjEzODE1MyIsImlhdCI6MTY1Njg4MTk1OSwic3ViIjoiQXV0aGVudGljYXRpb24iLCJ1c2VyX2lkIjoxMzgxNTMsInVzZXJuYW1lIjoiIiwicm9sZSI6InVzZXIiLCJpcF9hZGRyZXNzIjoiSVB2NDogMTI3LjAuMC4xICJ9.i8MGKafTqjlhvPs3xhxYNksco4Y01vxEJ7DKmQ7sIcs",
		},
	}

	for _, testCase := range testCases {
		res, err := authServiceInstance.ParseToken(testCase.sampleToken)

		if err != nil {
			t.Logf("%s", err.Error())
		} else {
			t.Logf("Result %v, Token Is Correct", res)
		}

	}
}

func TestGenerateUniqueSalt(t *testing.T) {
	rand.Seed(time.Now().Unix())

	for i := 0; i < 200; i++ {
		n := rand.Intn(randInstance)
		res := generateUniqueSalt(n)

		if len(res) < n || len(res) > n {
			t.Errorf("Incorrect Byte Generation")
			t.Fail()
		}
	}
}
