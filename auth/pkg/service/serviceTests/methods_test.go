package serviceTests

import (
	"auth/internal/entities"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type mockBehaviour func(s *MockAuthorizer, u entities.User)

	u := entities.User{
		Name:      "Nona",
		Surname:   "Davitashvili",
		UserName:  "nona.davitashvili",
		Email:     "nona.davitashvili@gmail.com",
		TelNumber: "+995592223344",
		IDNumber:  "11024029582",
		Password:  "Nonadavitashvili",
	}

	a := NewMockAuthorizer(t)

	d := func(s *MockAuthorizer, u entities.User) {
		a.EXPECT().CreateUser(u).Return(entities.AuthenticatedUserResponse{}, nil)
	}

}
