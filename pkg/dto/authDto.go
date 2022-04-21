package dto

import (
	"github.com/davitdarsalia/RestAPI.git/pkg/constants"
	"github.com/davitdarsalia/RestAPI.git/pkg/helpers"
	"github.com/davitdarsalia/RestAPI.git/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegUserDTO(u models.UserSignUp, c *gin.Context) {
	db := helpers.DbConnection()
	_, err := db.Exec(constants.UserSignUpQuery, u.Email, u.FirstName, u.LastName, u.Age, u.Password)
	if err != nil {
		c.String(http.StatusConflict, "User Already Exists")
		return
	}
	c.String(http.StatusCreated, "User Created")
	return
}

func CheckUserDTO(email, password string) (bool, string) {
	db := helpers.DbConnection()
	hash := helpers.HashData(password)
	password = hash
	var pass string
	var id string
	err := db.QueryRow(constants.CheckUser, password, email).Scan(&pass)
	if err != nil {
		return false, ""
	}
	err = db.QueryRow(constants.FetchUserId, email, password).Scan(&id)
	if err != nil {
		return false, ""
	}
	return true, id
}
