package service

import (
	"auth/internal/entities"
	"errors"
	"fmt"
	"log"
)

func (a *AuthService) CreateUser(u *entities.User) (entities.AuthenticatedUserResponse, error) {
	err := a.validator.Struct(u)
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
		fmt.Println(err)
		return entities.AuthenticatedUserResponse{}, err
	}

	aT, err := accessToken([]byte(u.Salt), id, a.credentials.TokenExpiryTime)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	rT, err := refreshToken([]byte(u.Salt), id)
	if err != nil {
		return entities.AuthenticatedUserResponse{}, err
	}

	return entities.AuthenticatedUserResponse{
		UserID:                id,
		AccessToken:           aT,
		AccessTokenExpiresAt:  fmt.Sprintf("%s Minutes", a.credentials.TokenExpiryTime),
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
		aT, err := accessToken([]byte(userInfo.Salt), userInfo.UserID, a.credentials.TokenExpiryTime)
		rT, err := refreshToken([]byte(userInfo.Salt), userInfo.UserID)

		if err != nil {
			log.Println("Token Generation Error")
		}

		return entities.AuthenticatedUserResponse{
			UserID:                userInfo.UserID,
			AccessToken:           aT,
			AccessTokenExpiresAt:  fmt.Sprintf("%s Minutes", a.credentials.TokenExpiryTime),
			RefreshToken:          rT,
			RefreshTokenExpiresAt: "13 days",
		}, nil
	}

	err = errors.New("user not found")

	return entities.AuthenticatedUserResponse{}, err
}

func (a *AuthService) RequestPasswordRecover(u *entities.RecoverPasswordInput) error {
	err := a.validator.Struct(u)
	if err != nil {
		return generateValidationStruct(err)
	}

	id, err := a.repo.RequestPasswordRecover(u)

	isUUID := checkUUID(id)

	fmt.Println(id)

	if err != nil || (err != nil && !isUUID) {
		return err
	}

	return nil
}

func (a *AuthService) ResetPassword(u *entities.RecoverPasswordInput) error {
	return nil
}
