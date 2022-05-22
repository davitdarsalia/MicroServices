package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"strconv"
)

func (a *AccountService) GetProfileDetails() (*entities.ProfileDetails, error) {
	id, _ := a.redisConn.Get(localContext, "UserID").Result()
	intID, _ := strconv.Atoi(id)

	return a.repo.GetProfileDetails(&intID)
}

func (a *AccountService) GetUserInfo() (*entities.UserInfo, error) {
	id, _ := a.redisConn.Get(localContext, "UserID").Result()
	intID, _ := strconv.Atoi(id)

	return a.repo.GetUserInfo(&intID)
}

func (a *AccountService) GetTrustedDevices() {
	//TODO implement me
}

func (a *AccountService) GetUserById() {
	//TODO implement me
}

func (a *AccountService) BlockUser() {
	//TODO implement me
}

func (a *AccountService) UnblockUser() {
	//TODO implement me
}

func (a *AccountService) BlockedUsersList() {
	//TODO implement me
}

func (a *AccountService) UploadProfileImage() {
	//TODO implement me
}

func (a *AccountService) LogoutSession() {
	//TODO implement me
}

func (a *AccountService) UpdateProfileDetails() {
	//TODO implement me
}

func (a *AccountService) UpdateTrustedDevices() {
	//TODO implement me
}

func (a *AccountService) SetPasscode() {
	//TODO implement me
}
