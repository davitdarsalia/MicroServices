package repository

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (r *AuthPostgres) RegisterUser(u *entities.User) (int, error) {
	var userId int
	query := fmt.Sprintf("INSERT INTO %s (personal_number, phonenum, username, email, firstname, lastname, ip_address, password, salt (UUID)) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING userid", "public.User")

	row := r.db.QueryRow(query, u.PersonalNumber, u.PhoneNumber, u.UserName, u.Email, u.FirstName, u.LastName, u.IpAddress, u.Password, u.Salt)

	if err := row.Scan(); err != nil {
		return 0, err
	}
	return userId, nil
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
