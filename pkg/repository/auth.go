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

func (r *AuthPostgres) CheckUser(username, password string) (entities.User, error) {
	var u entities.User

	err := r.db.Get(&u, constants.CheckUserQuery, username, password)

	return u, err
}

func (r *AuthPostgres) refreshLogin() {

}

func (r *AuthPostgres) resetPassword() {

}

func (r *AuthPostgres) resetPasswordProfile() {

}

func (r *AuthPostgres) otpGenerator() {

}
