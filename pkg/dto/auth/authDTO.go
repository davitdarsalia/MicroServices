package auth

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/handlers/security"
	"dbPractice/pkg/models"
	"net/http"
)

func CreateUserDTO(u models.UserSignUp, w http.ResponseWriter) (createUserErr bool) {
	dB := db.ConnectDB()
	_, err := dB.Exec(constants.UserSignUpQuery, u.Email, u.FirstName, u.LastName, u.Age, u.Password)
	if err != nil {
		createUserErr = true
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	return
}

func SignInUserDTO(email, password string) (bool, string) {
	dB := db.ConnectDB()
	hash := security.HashData(password)
	password = hash
	// Retrieved value based on Db query
	var passwordRes string
	var userId string
	err := dB.QueryRow(constants.CheckUser, password, email).Scan(&passwordRes)
	if err != nil {
		return false, ""
	}
	err = dB.QueryRow(constants.FetchUserId, email, password).Scan(&userId)
	if err != nil {
		return false, ""
	}
	return true, userId
}
