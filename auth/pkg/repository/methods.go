package repository

import (
	"context"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/queries"
	"github.com/davitdarsalia/auth/internal/utils"
)

func (d *DBInstance) Create(u *entities.User) (string, error) {
	var userID string

	err := d.db.QueryRow(context.Background(), queries.RegisterUserQuery,
		u.UserName, u.UserRole, u.Password, u.FirstName, u.LastName,
		u.Country, u.Email, u.Gender, u.City, u.CreatedAt, u.IpAddress,
	).Scan(&userID)

	utils.PgxErrorHandler(err)

	return userID, err
}

func (d *DBInstance) Login(email, password string) (string, error) {
	var userID string

	err := d.db.QueryRow(context.Background(), queries.CheckUserQuery, email, password).Scan(&userID)

	utils.PgxErrorHandler(err)

	return userID, err
}

func (d *DBInstance) Refresh() {
	//TODO implement me
}

func (d *DBInstance) Verify() {
	//TODO implement me
}

func (d *DBInstance) Reset() {
	//TODO implement me
}
