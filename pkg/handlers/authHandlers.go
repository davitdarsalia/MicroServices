package handlers

import (
	"encoding/json"
	"github.com/davitdarsalia/RestAPI.git/pkg/dto"
	"github.com/davitdarsalia/RestAPI.git/pkg/helpers"
	"github.com/davitdarsalia/RestAPI.git/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegistrationHandler(c *gin.Context) {
	var u models.UserSignUp
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	defer func() {
		closeErr := c.Request.Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	hashedPassword := helpers.HashData(u.Password)
	u.Password = hashedPassword
	dto.RegUserDTO(u, c)
	c.Header("Content-Type", "application-json")
}

func LoginHandler(c *gin.Context) {
	var u models.UserSignUp
	err := json.NewDecoder(c.Request.Body).Decode(&u)
	defer func() {
		closeErr := c.Request.Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()
	if err != nil {
		log.Fatal(err)
	}
	existence, userId := dto.CheckUserDTO(u.Email, u.Password)
	if existence == false {
		c.String(http.StatusNotFound, "User Does Not Exists")
		return
	}
	t := helpers.JWTGenerator(userId)
	bToken, marshalErr := json.Marshal(t)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	c.String(http.StatusOK, string(bToken))

}
func RefreshLoginHandler(c *gin.Context) {
	return
}

func LogoutHandler(c *gin.Context) {
	return
}
