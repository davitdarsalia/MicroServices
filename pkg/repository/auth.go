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

	row := r.db.QueryRow(constants.REGISTER_USER_QUERY, u.PersonalNumber, u.PhoneNumber, u.UserName, u.Email, u.FirstName, u.LastName, u.IpAddress, u.Password, u.Salt)

	if err := row.Scan(&userId); err != nil {
		return 0, err
	}

	return userId, nil
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
