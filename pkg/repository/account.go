package repository

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
)

func (r *AccountPostgres) GetProfileDetails(userid *int) (*entities.ProfileDetails, error) {
	var p entities.ProfileDetails

	r.db.QueryRow(constants.GetProfileDetails, userid).Scan(
		&p.ProfileImage,
		&p.Followers,
		&p.Following,
		&p.BlockedUsersAmount,
		&p.WorkingPlace,
		&p.Education,
		&p.Origin,
		&p.AdditionalEmail,
		&p.UserID,
	)

	return &p, nil
}

func (r *AccountPostgres) GetUserInfo() {
	//TODO implement me
}

func (r *AccountPostgres) GetTrustedDevices() {
	//TODO implement me
}

func (r *AccountPostgres) GetUserById() {
	//TODO implement me
}

func (r *AccountPostgres) BlockUser() {
	//TODO implement me
}

func (r *AccountPostgres) UnblockUser() {
	//TODO implement me
}

func (r *AccountPostgres) BlockedUsersList() {
	//TODO implement me
}

func (r *AccountPostgres) UploadProfileImage() {
	//TODO implement me
}

func (r *AccountPostgres) LogoutSession() {
	//TODO implement me
}

func (r *AccountPostgres) UpdateProfileDetails() {
	//TODO implement me
}

func (r *AccountPostgres) UpdateTrustedDevices() {
	//TODO implement me
}

func (r *AccountPostgres) SetPasscode() {
	//TODO implement me
}
