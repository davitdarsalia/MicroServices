package handler

import (
	"fmt"
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

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

func (h *Handler) BlockUser(c *gin.Context) {
	var b entities.BlockingUser

	if err := c.BindJSON(&b); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	err := h.services.BlockUser(&b)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.BlockUserError)
		return
	}

	c.JSON(http.StatusOK, entities.BlockUserResponse{
		Message:       constants.BlockUserSuccess,
		BlockedUserID: b.BlockedUserID,
	})

}
func (h *Handler) UnblockUser(c *gin.Context) {
	var b entities.UnblockingUser

	if err := c.BindJSON(&b); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	err := h.services.UnblockUser(&b)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.UnblockUserError)
		return
	}

	c.JSON(http.StatusOK, entities.BlockUserResponse{
		Message:       constants.UnblockUserSuccess,
		BlockedUserID: b.UnblockedUserID,
	})

}
func (h *Handler) BlockedUsersList(c *gin.Context) {
	userList, err := h.services.BlockedUsersList()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.GetBlockedUserListError)
		return
	}

	c.JSON(http.StatusOK, entities.BlockedUsersListResponse{
		Message:          constants.GetBlockedUserListSuccess,
		BlockedUsersList: userList,
		FetchedAt:        time.Now().Format(entities.RegularFormat),
	})

}
func (h *Handler) UploadProfileImage(c *gin.Context) {
	file, header, err := c.Request.FormFile(constants.ProfileImageFormFileHeader)

	// Checking For Nullable Header - File Size
	if header.Size == 0 || err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}
	err = h.services.UploadProfileImage(c, file)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.ProfileImageWithoutFileError)
		return
	}

	fmt.Println(file)

	uploadTime := time.Now().Format(entities.RegularFormat)

	c.JSON(http.StatusResetContent, entities.UploadProfileImageResponse{
		Message:    constants.ProfileImageUploadSuccess,
		UploadedAt: uploadTime,
	})

}
func (h *Handler) LogoutSession(c *gin.Context) {
	err := h.services.LogoutSession()
	fmt.Println(err)

}
func (h *Handler) UpdateProfileDetails(c *gin.Context) {

}

func (h *Handler) SetPasscode(c *gin.Context) {

}
