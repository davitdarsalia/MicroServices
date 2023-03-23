package serviceTests

import (
	"auth/internal/entities"
	"testing"
)

func BenchmarkCreateUser(b *testing.B) {
	u := entities.User{
		Name:      "Nona",
		Surname:   "Davitashvili",
		UserName:  "nona.davitashvili",
		Email:     "nona.davitashvili@gmail.com",
		TelNumber: "+995592223344",
		IDNumber:  "11024029582",
		Password:  "Nonadavitashvili",
	}

	a := NewMockAuthorizer(b)
	a.EXPECT().CreateUser(u).Return(entities.AuthenticatedUserResponse{}, nil)

	for i := 0; i < b.N; i++ {
		a.CreateUser(u)
	}
}

func BenchmarkLoginUser(b *testing.B) {
	u := entities.UserInput{
		Email:    "nona.davitashvili@gmail.com",
		IDNumber: "11024029582",
		Password: "Nonadavitashvili",
	}

	a := NewMockAuthorizer(b)
	a.EXPECT().LoginUser(u).Return(entities.AuthenticatedUserResponse{}, nil)

	for i := 0; i < b.N; i++ {
		a.LoginUser(u)
	}
}

func BenchmarkRecoveryPassword(b *testing.B) {
	u := entities.RecoverPasswordInput{
		Email:       "nona.davitashvili@gmail.com",
		IDNumber:    "11024029582",
		TelNumber:   "+995598299289",
		NewPassword: "Nona23151",
	}

	a := NewMockAuthorizer(b)
	a.EXPECT().RecoverPassword(u).Return(nil)

	for i := 0; i < b.N; i++ {
		a.RecoverPassword(u)
	}
}
