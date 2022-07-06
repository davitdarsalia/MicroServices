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

// @Summary Get Profile Details
// @Tags Account - Protected
// @Description Gets User Profile Details
// @ID user-profile-details
// @Accept json
// @Produce json
// @Success 200 {object} entities.GetProfileDetailsResponse
// @Failure 500 default {object} localError
// @Router /api/protected/account/profile-details [get]
func (h *Handler) getProfileDetails(c *gin.Context) {
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

// @Summary Get User Info
// @Tags Account - Protected
// @Description Get User Info
// @ID user-info
// @Accept json
// @Produce json
// @Success 200 {object} entities.GetUserInfoResponse
// @Failure 500 default {object} localError
// @Router /api/protected/account/user-info [get]
func (h *Handler) getUserInfo(c *gin.Context) {
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

// @Summary Get Trusted Devices
// @Tags Account - Protected
// @Description Get Trusted Devices List
// @ID trusted-device
// @Accept json
// @Produce json
// @Success 200 {object} entities.GetTrustedDevices
// @Failure 500 default {object} localError
// @Router /api/protected/account/trusted-devices-list [get]
func (h *Handler) getTrustedDevices(c *gin.Context) {
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

// @Summary Add Trusted Device
// @Security ApiKeyAuth
// @Tags Account - Protected
// @Description Adds Trusted Devices (Device IP)
// @ID add-trusted-device
// @Accept json
// @Produce json
// @Param input body entities.TrustedDevices true "Credentials"
// @Success 205 {object} entities.TrustedDevicesResponse
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/add-trusted-device [post]
func (h *Handler) addTrustedDevice(c *gin.Context) {
	// TODO - Make Ip Unique For DBMS
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

// @Summary Block User
// @Security ApiKeyAuth
// @Tags Account - Protected
// @Description Block User By ID
// @ID block-user
// @Accept json
// @Produce json
// @Param input body entities.BlockingUser true "Credentials"
// @Success 200 {object} entities.BlockUserResponse
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/block-user [post]
func (h *Handler) blockUser(c *gin.Context) {
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

// @Summary Unblock User
// @Security ApiKeyAuth
// @Tags Account - Protected
// @Description Unblock Blocked User By ID
// @ID unblock-user
// @Accept json
// @Produce json
// @Param input body entities.UnblockingUser true "Credentials"
// @Success 200 {object} entities.UnblockUserResponse
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/unblock-user [post]
func (h *Handler) unblockUser(c *gin.Context) {
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

	c.JSON(http.StatusOK, entities.UnblockUserResponse{
		Message:         constants.UnblockUserSuccess,
		UnblockedUserID: b.UnblockedUserID,
	})

}

// @Summary Get Blocked Users List
// @Tags Account - Protected
// @Description Gets Blocked Users List
// @ID blocked-users-list
// @Accept json
// @Produce json
// @Success 200 {object} entities.BlockedUsersListResponse
// @Failure 500 default {object} localError
// @Router /api/protected/account/blocked-user-list [get]
func (h *Handler) blockedUsersList(c *gin.Context) {
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

// @Summary Upload Profile Image
// @Security ApiKeyAuth
// @Tags Account - Protected
// @Description Add Profile Image (Uploading Multiple Times Is Acceptable)
// @ID upload-profile-image
// @Accept json
// @Produce json
// @Param input body entities.UploadProfileImageResponse true "Credentials"
// @Success 200 {object} entities.UploadProfileImageResponse
// @Failure 400 {object} localError
// @Failure 500 default {object} localError
// @Router /api/auth/unblock-user [post]
func (h *Handler) uploadProfileImage(c *gin.Context) {
	file, header, err := c.Request.FormFile(constants.ProfileImageFormFileHeader)

	if header.Size == 0 || err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.BadRequest)
		return
	}

	uploadTime := time.Now().Format(entities.RegularFormat)

	err = h.services.UploadProfileImage(c, file, &uploadTime)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.ProfileImageWithoutFileError)
		return
	}

	c.JSON(http.StatusResetContent, entities.UploadProfileImageResponse{
		Message:    constants.ProfileImageUploadSuccess,
		UploadedAt: uploadTime,
	})

}

// @Summary Get Images
// @Tags Account - Protected
// @Description Get Uploaded Images
// @ID uploaded-images
// @Accept json
// @Produce json
// @Success 200 {object} entities.GetImagesResponse
// @Failure 500 default {object} localError
// @Router /api/protected/account/get-images [get]
func (h *Handler) getImages(c *gin.Context) {
	images, err := h.services.GetImages()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, constants.GetImageErrors)
		return
	}

	c.JSON(http.StatusOK, entities.GetImagesResponse{
		Message: constants.GetImageSuccess,
		Images:  images,
	})

}

func (h *Handler) logoutSession(c *gin.Context) {
	err := h.services.LogoutSession()
	fmt.Println(err)

}
func (h *Handler) updateProfileDetails(c *gin.Context) {

}

func (h *Handler) setPasscode(c *gin.Context) {

}
