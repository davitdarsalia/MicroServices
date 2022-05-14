package repository

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) RegisterUser(u *entities.User) (int, error) {
	var userId int

	row := r.db.QueryRow(constants.RegisterUserQuery, u.PersonalNumber, u.PhoneNumber, u.UserName, u.Email, u.FirstName, u.LastName, u.IpAddress, u.Password, u.Salt)

	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *AuthPostgres) CheckUser(username, password string) (int, error) {
	var userID int

	row := r.db.QueryRow(constants.CheckUserQuery, username, password)

	if err := row.Scan(&userID); err != nil {
		return 0, nil
	}

	return userID, nil
}
