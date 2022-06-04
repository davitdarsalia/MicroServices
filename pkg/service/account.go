package service

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"log"
	"strconv"
)

func (a *AccountService) GetProfileDetails() (*entities.ProfileDetails, error) {
	id, err := a.redisConn.Get(localContext, "UserID").Result()

	if err != nil {
		log.Fatalf("%s : %s", err, "RedisGetError")
	}
	intID, _ := strconv.Atoi(id)

	return a.repo.GetProfileDetails(&intID)
}

func (a *AccountService) GetUserInfo() (*entities.UserInfo, error) {
	id, err := a.redisConn.Get(localContext, constants.RedisID).Result()
	intID, _ := strconv.Atoi(id)

	if err != nil {
		log.Fatalf("%s : %s", err, "RedisGetError")
	}

	return a.repo.GetUserInfo(&intID)
}

func (a *AccountService) GetTrustedDevices() ([]entities.TrustedDevices, error) {
	id, err := a.redisConn.Get(localContext, constants.RedisID).Result()
	intID, _ := strconv.Atoi(id)

	if err != nil {
		log.Fatalf("%s : %s", err, "RedisGetError")
	}

	return a.repo.GetTrustedDevices(&intID)
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

func (a *AccountService) AddTrustedDevice(r *entities.TrustedDevices) (int, error) {
	id, err := a.redisConn.Get(localContext, "UserID").Result()

	r.DeviceID = generateRandNumber(1, 1000000)
	r.DeviceIpAddress = entities.GetIp()
	r.DeviceIdentifier = generateUniqueSalt(20)

	fmt.Println(r.DeviceIdentifier)

	if err != nil {
		log.Fatal(err)
	}

	intID, _ := strconv.Atoi(id)

	return a.repo.AddTrustedDevice(&intID, r)
}

func (a *AccountService) SetPasscode() {
	//TODO implement me
}
