package repository

import (
	"auth/internal/entities"
	"auth/internal/queries"
	"context"
)

func (a *AuthPostgres) CreateUser(u *entities.User) (string, error) {
	var userID string

	row := a.db.QueryRow(context.Background(), queries.CreateUserQuery,
		u.Name, u.Surname, u.UserName, u.Email, u.TelNumber,
		u.IDNumber, u.Password, u.DateCreated, u.IPAddress, u.Salt,
	)

	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (a *AuthPostgres) LoginUser(u entities.UserInput) (entities.UserMetaInfo, error) {
	var userID string
	var salt string
	var password string

	row := a.db.QueryRow(context.Background(), queries.LoginUserQuery, u.Email, u.IDNumber)

	if err := row.Scan(&userID, &salt, &password); err != nil {
		return entities.UserMetaInfo{}, err
	}

	return entities.UserMetaInfo{
		Password: password,
		Salt:     salt,
		UserID:   userID,
	}, nil
}

func (a *AuthPostgres) RequestPasswordRecover(u *entities.RecoverPasswordInput) (string, error) {
	var userID string

	row := a.db.QueryRow(context.Background(), queries.UpdatePasswordQuery, u.Email, u.IDNumber, u.TelNumber)

	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (a *AuthPostgres) ResetPassword(u *entities.RecoverPasswordInput) (string, error) {
	//_, err := a.db.Exec(queries.UpdatePasswordQuery, u.NewPassword, u.Email, u.IDNumber, u.TelNumber)
	//
	//if err != nil {
	//	return err
	//}
	return "", nil
}
