package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
	"strconv"
)

func (a *AccountService) GetProfileDetails() (*entities.ProfileDetails, error) {
	id, _ := a.redisConn.Get(localContext, "UserID").Result()
	intID, err := strconv.Atoi(id)

	if err != nil {
		log.Println("[Account Service] - ParseInt Handler Error")
	}

	return a.repo.GetProfileDetails(intID)
}

func (a *AccountService) GetUserInfo() {
	//TODO implement me
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
