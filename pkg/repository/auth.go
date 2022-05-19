package repository

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"time"
)

func (r *AuthPostgres) RegisterUser(u *entities.User) (int, error) {
	var userId int
	err := r.db.QueryRow(constants.RegisterUserQuery, u.PersonalNumber, u.PhoneNumber, u.UserName, u.Email, u.FirstName, u.LastName, u.IpAddress, u.Password, u.Salt).Scan(&userId)

	return userId, err
}

func (r *AuthPostgres) CheckUser(username, password string) (entities.User, error) {
	var u entities.User
	err := r.db.Get(&u, constants.CheckUserQuery, username, password)

	return u, err
}

func (r *AuthPostgres) ResetPassword(p *entities.ResetPassword) (string, error) {
	var userID string
	err := r.db.Get(&userID, constants.CheckUserByEmail, p.Email, p.UserName, p.PersonalNumber)

	if err != nil {
		return "", err
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

func (r *AuthPostgres) RefreshLogin() {

}
