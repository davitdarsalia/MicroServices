package service

import (
	"auth/internal/entities"
	"errors"
	"fmt"
	"log"
	"os"
)

func (a *AuthService) CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error) {
	err := a.validator.Struct(&u)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, generateValidationStruct(err)
	}

	s, err := salt()
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}
	u.Salt = string(s)

	hashedPass, err := hash(u.Password, u.Salt)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	u.Password = hashedPass
	u.IPAddress = getIPv6()
	u.DateCreated = getFormattedDateTime()

	id, err := a.repo.CreateUser(u)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	aT, err := accessToken([]byte(u.Salt), id)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	rT, err := refreshToken()
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	return entities.AuthenticatedUserResponse{
		UserID:                id,
		AccessToken:           aT,
		AccessTokenExpiresAt:  fmt.Sprintf("%s Minutes", os.Getenv("TOKEN_EXPIRY_TIME")),
		RefreshToken:          rT,
		RefreshTokenExpiresAt: "13 days",
	}, nil
}

func (a *AuthService) LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error) {
	err := a.validator.Struct(&u)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, generateValidationStruct(err)
	}

	userInfo, err := a.repo.LoginUser(u)

	hashedPass, err := hash(u.Password, userInfo.Salt)

	if userInfo.Password == hashedPass {
		aT, err := accessToken([]byte(userInfo.Salt), userInfo.UserID)
		rT, err := refreshToken()

		if err != nil {
			log.Println("Token Generation Error")
		}

		return entities.AuthenticatedUserResponse{
			UserID:                userInfo.UserID,
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

	//publisher := a.messageQueue

	if err != nil {
		return generateValidationStruct(err)
	}

	//newSalt, err := generateSalt()
	//
	//if err != nil {
	//	log.Printf("Salt Generation Error, %s", err.Error())
	//}
	//
	//u.NewPassword = hash(u.NewPassword, newSalt)

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
