package repository

import (
	"time"

	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
)

func (r *AuthPostgres) RegisterUser(u *entities.User) (int, error) {
	var userId int
	err := r.db.QueryRow(constants.RegisterUserQuery, u.PersonalNumber, u.PhoneNumber, u.UserName, u.Email, u.FirstName, u.LastName, u.IpAddress, u.Password, u.Salt).Scan(&userId)

	// Setting Default Values For Some Tables In Order To Update Them Later (Or Not)

	if err == nil {
		// If We've Successfully Registered An User, Then Default Values In Settings Will Be Applied
		go func() {
			r.db.Query(constants.InitNotificationSettings, userId)
		}()

		go func() {
			r.db.Query(constants.InitPaymentSettings, userId)
		}()

		go func() {
			r.db.Query(constants.InitSecuritySettings, userId)
		}()

		// TODO - Add Privacy Settings Query
	}

	return userId, err
}

func (r *AuthPostgres) CheckUser(username, password string) (entities.User, error) {
	c := make(chan error, 1)
	var u entities.User

	go func() {
		c <- r.db.Get(&u, constants.CheckUserQuery, username, password)
	}()

	defer close(c)
	return u, <-c
}

func (r *AuthPostgres) ResetPassword(p *entities.ResetPassword) (string, error) {
	var userID string
	c := make(chan error, 1)

	go func() {
		c <- r.db.Get(&userID, constants.CheckUserByEmail, p.Email, p.UserName, p.PersonalNumber)
	}()

	if <-c != nil {
		return "", <-c
	}

	return userID, nil
}

func (r *AuthPostgres) ValidateResetEmail(p *entities.ValidateResetEmail) error {
	// TODO - Implement Get Location Function (Lat. Lng. City, Country Is Optional)

	_, err := r.db.Exec(constants.InsertProfileActivityResetPassword, p.PersonalNumber, time.Now(), time.Now(), "Tbilisi")
	_, err = r.db.Exec(constants.UpdatePassword, p.NewPassword, p.PersonalNumber)

	return err
}

func (r *AuthPostgres) ResetPasswordProfile(e *entities.ResetPasswordInput) error {
	_, err := r.db.Exec(constants.InsertProfileActivityResetPasswordProfile, time.Now(), time.Now(), "Tbilisi")
	_, err = r.db.Exec(constants.UpdatePasswordFromProfile, e.NewPassword, e.UserName)

	return err
}
