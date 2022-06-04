package repository

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
)

func (r *AccountPostgres) GetProfileDetails(userID *int) (*entities.ProfileDetails, error) {
	var p entities.ProfileDetails

	r.db.QueryRow(constants.GetProfileDetails, userID).Scan(
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

func (r *AccountPostgres) GetUserInfo(userID *int) (*entities.UserInfo, error) {
	var p entities.UserInfo

	r.db.QueryRow(constants.GetUserInfo, userID).Scan(
		&p.UserID,
		&p.PersonalNumber,
		&p.PhoneNumber,
		&p.UserName,
		&p.Email,
		&p.FirstName,
		&p.LastName,
		&p.IpAddress,
		&p.Password,
		&p.Salt,
		&p.ProfileImage,
		&p.Followers,
		&p.Following,
		&p.BlockedUsersAmount,
		&p.WorkingPlace,
		&p.Education,
		&p.Origin,
		&p.AdditionalEmail,
		&p.Education,
	)

	return nil, nil
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

func (r *AccountPostgres) AddTrustedDevice(userID *int, t *entities.TrustedDevices) (int, error) {
	_, err := r.db.Exec(constants.AddTrustedDevice, *userID,
		t.DeviceID, t.DeviceName, t.DeviceIpAddress, t.DeviceIdentifier)

	return *userID, err
}

func (r *AccountPostgres) SetPasscode() {
	//TODO implement me
}
