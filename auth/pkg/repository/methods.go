package repository

import (
	"auth/internal/entities"
	"auth/internal/queries"
)

func (a *AuthPostgres) CreateUser(u entities.User) (string, error) {
	var userID string

	row := a.db.QueryRow(queries.CreateUserQuery, u.Name, u.Surname, u.UserName, u.Email, u.TelNumber, u.IDNumber, u.Password, u.DateCreated, u.IPAddress, u.Salt)

	if err := row.Scan(&userID); err != nil {
		return "No User Found", err
	}

	return userID, nil
}

func (a *AuthPostgres) LoginUser(u entities.UserInput) ([3]string, error) {
	var userID string
	var salt string
	var password string

	row := a.db.QueryRow(queries.LoginUserQuery, u.Email, u.IDNumber)

	if err := row.Scan(&userID, &salt, &password); err != nil {
		return [3]string{"", "", ""}, err
	}

	var userData = [3]string{password, salt, userID}

	return userData, nil
}

func (a *AuthPostgres) RecoverPassword(u entities.RecoverPasswordInput) error {
	_, err := a.db.Exec(queries.UpdatePasswordQuery, u.NewPassword, u.Email, u.IDNumber, u.TelNumber)

	if err != nil {
		return err
	}
	return nil
}
