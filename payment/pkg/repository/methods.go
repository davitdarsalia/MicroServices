package repository

import (
	"context"
	"fmt"
	"github.com/davitdarsalia/payment/internal/entities"
	"github.com/davitdarsalia/payment/internal/queries"
	"github.com/davitdarsalia/payment/internal/types"
	"github.com/davitdarsalia/payment/internal/utils"
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

func (d *DBInstance) Reset(email, idNumber string, newPassword types.Hash512) error {
	var userID string

	fmt.Println(email, idNumber, newPassword, "DDD")

	err := d.db.QueryRow(context.Background(), queries.CheckUserForReset, email, idNumber).Scan(&userID)

	if err != nil {
		return err
	}

	_, err = d.db.Exec(context.Background(), queries.RegisterUserQuery, newPassword, userID)

	if err != nil {
		utils.PgxErrorHandler(err)
		return err
	}

	return nil
}

func (d *DBInstance) Verify() {
	//TODO implement me
}
