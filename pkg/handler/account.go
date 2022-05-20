package handler

import (
	"github.com/davitdarsalia/LendAppBackend/constants"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/gin-gonic/gin"
	"net/http"
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

}
func (h *Handler) GetTrustedDevices(c *gin.Context) {

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
func (h *Handler) UpdateTrustedDevices(c *gin.Context) {

}
func (h *Handler) SetPasscode(c *gin.Context) {

}
