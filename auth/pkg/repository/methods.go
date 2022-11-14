package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/davitdarsalia/auth/internal/entities"
	"github.com/davitdarsalia/auth/internal/queries"
	"github.com/jackc/pgx/v5/pgconn"
)

func (d DBInstance) Create(u *entities.User) (string, error) {
	var userID string

	err := d.db.QueryRow(context.Background(), queries.RegisterUserQuery,
		u.UserName, u.UserRole, u.Password, u.FirstName, u.LastName,
		u.Country, u.Email, u.Gender, u.City, u.CreatedAt, u.IpAddress,
	).Scan(&userID)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			fmt.Println(pgErr.Message) // => syntax error at end of input
			fmt.Println(pgErr.Code)    // => 42601
		}
	}

	return userID, err
}

func (d DBInstance) Login() {
	//TODO implement me
	panic("implement me")
}

func (d DBInstance) Refresh() {
	//TODO implement me
	panic("implement me")
}

func (d DBInstance) Verify() {
	//TODO implement me
	panic("implement me")
}

func (d DBInstance) Reset() {
	//TODO implement me
	panic("implement me")
}
