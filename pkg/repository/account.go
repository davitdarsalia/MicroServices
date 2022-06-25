package repository

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
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

	p.Password = ""

	return &p, nil
}

func (r *AccountPostgres) GetTrustedDevices(userID *int) ([]entities.TrustedDevices, error) {
	var devices []entities.TrustedDevices

	deviceRows, err := r.db.Query(constants.GetTrustedDevices, userID)
	defer deviceRows.Close()

	if err != nil {
		log.Println(err, constants.GetTrustedDevicesError)
	}

	for deviceRows.Next() {
		var deviceInstance entities.TrustedDevices
		err := deviceRows.Scan(
			&deviceInstance.UserID,
			&deviceInstance.DeviceID,
			&deviceInstance.DeviceName,
			&deviceInstance.DeviceIpAddress,
			&deviceInstance.DeviceIdentifier,
		)
		if err != nil {
			log.Printf("%s: %s", err, "Get Trusted Devices List Repository - Error During Scanning Rows")
		}
		devices = append(devices, deviceInstance)
	}

	return devices, nil
}

func (r *AccountPostgres) BlockUser(userID *int, u *entities.BlockingUser) error {
	_, err := r.db.Exec(constants.BlockUserQuery, userID, u.BlockedUserID, u.BlockedAt)

	return err
}

func (r *AccountPostgres) UnblockUser(userID *int, u *entities.UnblockingUser) error {
	_, err := r.db.Exec(constants.UnblockUserQuery, userID, u.UnblockedUserID)

	return err
}

func (r *AccountPostgres) BlockedUsersList(userID *int) ([]entities.BlockedUsersList, error) {
	var l []entities.BlockedUsersList

	blockedUsersRows, err := r.db.Query(constants.GetBlockedUsersList, userID)
	defer blockedUsersRows.Close()

	if err != nil {
		log.Println(err, constants.GetBlockedUserListError)
	}

	for blockedUsersRows.Next() {
		var u entities.BlockedUsersList

		err := blockedUsersRows.Scan(
			&u.UserID,
			&u.BlockedUserID,
			&u.BlockedAt,
		)

		if err != nil {
			log.Printf("%s: %s", err, "Get Blocked Users List Repository - Error During Scanning Rows")
		}

		l = append(l, u)
	}

	return l, nil
}

func (r *AccountPostgres) UploadProfileImage(f string, userID int, uploadTime *string) error {
	_, err := r.db.Exec(constants.AddProfileImage, userID, f, *uploadTime)

	return err
}

func (r *AccountPostgres) GetImages(userID *int) ([]entities.Image, error) {
	var images []entities.Image

	imageRows, err := r.db.Query(constants.GetImages, userID)

	defer imageRows.Close()

	if err != nil {
		log.Println(err, constants.GetImageErrors)
	}

	for imageRows.Next() {
		var imageInstance entities.Image
		err := imageRows.Scan(
			&imageInstance.Image,
			&imageInstance.UploadedAt,
			&imageInstance.ImageID,
			&imageInstance.IsProfileImage,
		)

		if err != nil {
			log.Printf("%s: %s", err, "Get Images List Repository - Error During Scanning Rows")
		}
		images = append(images, imageInstance)
	}

	return images, nil
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
