package service

import (
	"auth/internal/entities"
	"errors"
	"log"
)

/* Methods */

func (a *AuthService) CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error) {
	err := a.validator.Struct(&u)

	if err != nil {
		return entities.AuthenticatedUserResponse{}, generateValidationStruct(err)
	}

	salt, _ := generateSalt()

	u.Salt = salt
	u.Password = hash(u.Password, u.Salt)
	u.IPAddress = getIPv6()
	u.DateCreated = getFormattedDateTime()

	id, err := a.repo.CreateUser(u)

	if err == nil {
		aT, err := accessToken(id, u.Salt)
		rT, err := refreshToken(id, u.Salt)

		if err != nil {
			log.Println("Token Generation Error")
		}

		return entities.AuthenticatedUserResponse{
			UserID:                id,
			AccessToken:           aT,
			AccessTokenExpiresAt:  "200 Minutes",
			RefreshToken:          rT,
			RefreshTokenExpiresAt: "200 Hours",
		}, nil
	}

	return entities.AuthenticatedUserResponse{
		AccessToken: "No access",
	}, err

}

func (a *AuthService) LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error) {
	err := a.validator.Struct(&u)

	if err != nil {
		return entities.AuthenticatedUserResponse{}, generateValidationStruct(err)
	}

	data, err := a.repo.LoginUser(u)

	if data[0] == hash(u.Password, data[1]) {
		aT, err := accessToken(data[2], data[1])
		rT, err := refreshToken(data[2], data[1])

		if err != nil {
			log.Println("Token Generation Error")
		}

		return entities.AuthenticatedUserResponse{
			UserID:                data[2],
			AccessToken:           aT,
			AccessTokenExpiresAt:  "200 Minutes",
			RefreshToken:          rT,
			RefreshTokenExpiresAt: "200 Hours",
		}, nil
	}

	err = errors.New("user not found")

	return entities.AuthenticatedUserResponse{}, err
}

func (a *AuthService) RecoverPassword(u entities.RecoverPasswordInput) error {
	err := a.validator.Struct(&u)

	if err != nil {
		return generateValidationStruct(err)
	}

	newSalt, err := generateSalt()

	if err != nil {
		log.Printf("Salt Generation Error, %s", err.Error())
	}

	u.NewPassword = hash(u.NewPassword, newSalt)

	// Add code verification
	return a.repo.RecoverPassword(u)
}

func (a *AuthService) CheckToken(authToken, signKey string) (string, error) {
	userID, err := checkToken(authToken, signKey)

	if err != nil {
		return "Not Authorized", err
	}

	return userID, nil
}
