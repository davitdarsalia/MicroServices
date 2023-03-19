package repository

import (
	"auth/internal/entities"
	"auth/internal/queries"
	"context"
)

const (
	CreateUserTransactionStatement = "create-user"
)

func (a *AuthPostgres) CreateUser(u entities.User) (string, error) {
	var userID string

	tx, err := a.db.Begin(context.Background())
	if err != nil {
		return "", err
	}

	row := tx.QueryRow(context.Background(), queries.CreateUserQuery,
		u.Name, u.Surname, u.UserName, u.Email, u.TelNumber,
		u.IDNumber, u.Password, u.DateCreated, u.IPAddress, u.Salt,
	)

	err = row.Scan(&userID)

	if err != nil {
		return "", err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return "", err
	}

	//ps, err := a.db.Prepare(context.Background(), "create-user", stmt)
	//if err != nil {
	//	return "", fmt.Errorf("failed to prepare create user statement: %v", err)
	//}
	//defer ps.Close()
	//
	//// Execute the statement and retrieve the new user's ID
	//if _, err := ps.Exec(context.Background(), u.Name, u.Surname, u.UserName, u.Email, u.TelNumber, u.IDNumber, u.Password, u.DateCreated, u.IPAddress, u.Salt).Scan(&userID); err != nil {
	//	if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // Check for unique constraint violation error
	//		if strings.Contains(pgErr.Detail, "email") {
	//			return "", fmt.Errorf("email already exists")
	//		} else if strings.Contains(pgErr.Detail, "username") {
	//			return "", fmt.Errorf("username already exists")
	//		} else {
	//			return "", fmt.Errorf("failed to create user: %v", err)
	//		}
	//	}
	//	return "", fmt.Errorf("failed to create user: %v", err)
	//}

	// Learn pgx transactions

	return userID, nil

}

func (a *AuthPostgres) LoginUser(u entities.UserInput) ([3]string, error) {
	//var userID string
	//var salt string
	//var password string

	//row := a.db.QueryRow(queries.LoginUserQuery, u.Email, u.IDNumber)
	//
	//if err := row.Scan(&userID, &salt, &password); err != nil {
	//	return [3]string{"", "", ""}, err
	//}
	//
	//var userData = [3]string{password, salt, userID}

	return [3]string{}, nil
}

func (a *AuthPostgres) RecoverPassword(u entities.RecoverPasswordInput) error {
	//_, err := a.db.Exec(queries.UpdatePasswordQuery, u.NewPassword, u.Email, u.IDNumber, u.TelNumber)
	//
	//if err != nil {
	//	return err
	//}
	return nil
}
