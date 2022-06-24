package service

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"
)

func (a *AccountService) GetProfileDetails() (*entities.ProfileDetails, error) {
	id := a.getRedisUserID()

	return a.repo.GetProfileDetails(&id)
}

func (a *AccountService) GetUserInfo() (*entities.UserInfo, error) {
	id := a.getRedisUserID()

	return a.repo.GetUserInfo(&id)
}

func (a *AccountService) GetTrustedDevices() ([]entities.TrustedDevices, error) {
	id := a.getRedisUserID()

	return a.repo.GetTrustedDevices(&id)
}

func (a *AccountService) BlockUser(b *entities.BlockingUser) error {
	id := a.getRedisUserID()

	b.BlockedAt = formatNowDate()

	return a.repo.BlockUser(&id, b)
}

func (a *AccountService) UnblockUser(b *entities.UnblockingUser) error {
	id := a.getRedisUserID()

	return a.repo.UnblockUser(&id, b)

}

func (a *AccountService) BlockedUsersList() ([]entities.BlockedUsersList, error) {
	id := a.getRedisUserID()
	fmt.Println(id)

	return a.repo.BlockedUsersList(&id)
}

func (a *AccountService) UploadProfileImage(c *gin.Context, f multipart.File, uploadTime *string) error {
	id := a.getRedisUserID()

	reader := bufio.NewReader(f)
	contentBytes, _ := ioutil.ReadAll(reader)

	encodedProfileImage := base64.StdEncoding.EncodeToString(contentBytes)

	return a.repo.UploadProfileImage(encodedProfileImage, id, uploadTime)
}

func (a *AccountService) LogoutSession() error {
	a.redisConn.Append(localContext, constants.SessionID, constants.NotValidMark)
	s, _ := a.redisConn.Get(localContext, constants.SessionID).Result()

	fmt.Println(s)
	return nil
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
