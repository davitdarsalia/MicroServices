package handler

import (
	"auth/internal/entities"
	"auth/pkg/service"
	mock "auth/pkg/service/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/quick"
)

type statusCode int
type uuid string

func checkUUID(uuid uuid) bool {
	return uuid != "" && len(uuid) == 36 && uuid[14] == '4' && (uuid[19] == '8' || uuid[19] == '9' || uuid[19] == 'a' || uuid[19] == 'b') && uuid[8] == '-' && uuid[13] == '-' && uuid[18] == '-' && uuid[23] == '-'
}

func TestHandlerSignUp(t *testing.T) {
	type mockBehavior func(s *mock.MockAuthorizer, u entities.User)

	type userTest struct {
		testID             string
		inputBody          string
		inputUser          entities.User
		mockBehavior       mockBehavior
		expectedStatusCode statusCode
	}

	testTable := []userTest{
		{
			testID:    "OK",
			inputBody: `{"name": "David","surname": "Darsalia","username": "David.1998!", "email": "darsalia.david1998@gmail.com", "tel_number": "+995598299289","id_number": "01027090407","password": "David.1998!"}`,
			inputUser: entities.User{
				Name:      "David",
				Surname:   "Darsalia",
				UserName:  "David.1998!",
				Email:     "darsalia.david1998@gmail.com",
				TelNumber: "+995598299289",
				IDNumber:  "01027090407",
				Password:  "David.1998!",
			},
			mockBehavior: func(s *mock.MockAuthorizer, u entities.User) {
				s.EXPECT().CreateUser(u).Return([3]string{"", "", ""})

			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, c := range testTable {
		t.Run(c.testID, func(t *testing.T) {
			endpoint := "/sign-up"
			// For every testCase, create a controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Mock interface
			authorizer := mock.NewMockAuthorizer(ctrl)
			c.mockBehavior(authorizer, c.inputUser)

			service := &service.Service{Authorizer: authorizer}
			handler := New(service)

			router := gin.Default()
			router.POST(endpoint, handler.createUser)

			// Generating test request - create responseWriter and request
			w := httptest.NewRecorder()
			request := httptest.NewRequest("POST", endpoint, bytes.NewBufferString(c.inputBody))

			// Perform request

			router.ServeHTTP(w, request)

			assert.Equal(t, c.expectedStatusCode, w.Code)

			// Test uuid
			valid := checkUUID("4fc42d9b-5566-4ea7-9bef-07c9556b23c7")
			exampleUUID := "4fc42d9b-5566-4ea7-9bef-07c9556b23c7"
			if err := quick.Check(valid, &quick.Config{
				MaxCount: len(exampleUUID),
			}); err != nil {
				t.Error(err)
			}

		})
	}
}
