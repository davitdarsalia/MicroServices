package dto

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
)

func CreateUserDTO(u models.UserSignUp) {
	dB := db.ConnectDB()

	_, err := dB.Exec(constants.UserSignUpQuery, u.Email, u.FirstName, u.LastName, u.Age, u.Password)

	if err != nil {
		log.Fatal(err)
	}
}

func SignInUserDTO(email, password string) bool {
	dB := db.ConnectDB()

	var emailRes string
	var passwordRes string

	err := dB.QueryRow(constants.CheckUser, email, password).Scan(&emailRes, &passwordRes)

	if err != nil {
		return false
	}

	return true

}
