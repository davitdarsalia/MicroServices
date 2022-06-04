package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UploadProfilePicture(c *gin.Context) {
	//file, handler, err := c.Request.FormFile("profilePic")
	//
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//defer file.Close()
	//
	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("Size: %+v\n", handler.Size)
	//fmt.Printf("Mime Header: %+v\n", handler.Header)
	//
	//tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer tempFile.Close()
	//
	////fileBytes, err := ioutil.ReadAll(file)

}

func (h *Handler) GetProfileDetails(c *gin.Context) {
	p, err := h.services.GetProfileDetails()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.GetProfileDetailsError)
		return
	}

	c.JSON(http.StatusOK, entities.GetProfileDetailsResponse{
		Message:            constants.GetProfileDetailsSuccess,
		ProfileImage:       p.ProfileImage,
		Followers:          p.Followers,
		Following:          p.Following,
		BlockedUsersAmount: p.BlockedUsersAmount,
		WorkingPlace:       p.WorkingPlace,
		Education:          p.Education,
		Origin:             p.Origin,
		AdditionalEmail:    p.AdditionalEmail,
		UserID:             p.UserID,
	})

}
func (h *Handler) GetUserInfo(c *gin.Context) {
	p, err := h.services.GetUserInfo()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.GetUserInfoError)
		return
	}

	id, _ := strconv.Atoi(p.UserName)

	c.JSON(http.StatusOK, entities.GetUserInfoResponse{
		Message: constants.GetUserInfoSuccess,
		User: entities.User{
			UserID:         id,
			PersonalNumber: p.PersonalNumber,
			PhoneNumber:    p.PhoneNumber,
			UserName:       p.UserName,
			Email:          p.Email,
			FirstName:      p.FirstName,
			LastName:       p.LastName,
			IpAddress:      p.IpAddress,
			Password:       p.Password,
			Salt:           p.Salt,
		},
		ProfileImage:       p.ProfileImage,
		Followers:          p.Followers,
		Following:          p.Following,
		BlockedUsersAmount: p.BlockedUsersAmount,
		WorkingPlace:       p.WorkingPlace,
		Education:          p.Education,
		Origin:             p.Origin,
		AdditionalEmail:    p.AdditionalEmail,
	})

}
func (h *Handler) GetTrustedDevices(c *gin.Context) {
	p, err := h.services.GetTrustedDevices()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.GetTrustedDevicesError)
		return
	}

	c.JSON(http.StatusOK, entities.GetTrustedDevices{
		Message:    constants.GetTrustedDevicesSuccess,
		DeviceList: p,
	})
}

// AddTrustedDevice TODO - Make Ip Unique For DBMS
func (h *Handler) AddTrustedDevice(c *gin.Context) {
	var d entities.TrustedDevices

	if err := c.BindJSON(&d); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	id, err := h.services.Account.AddTrustedDevice(&d)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.AddTrustedDeviceError)
		return
	}

	c.JSON(http.StatusResetContent, entities.TrustedDevicesResponse{
		Message: constants.AddTrustedDeviceSuccess,
		UserID:  fmt.Sprintf("%d", id),
	})

}

func (h *Handler) GetUserById(c *gin.Context) {

}
func (h *Handler) BlockUser(c *gin.Context) {

}
func (h *Handler) UnblockUser(c *gin.Context) {

}
func (h *Handler) BlockedUsersList(c *gin.Context) {

}
func (h *Handler) UploadProfileImage(c *gin.Context) {

}
func (h *Handler) LogoutSession(c *gin.Context) {

}
func (h *Handler) UpdateProfileDetails(c *gin.Context) {

}

func (h *Handler) SetPasscode(c *gin.Context) {

}
