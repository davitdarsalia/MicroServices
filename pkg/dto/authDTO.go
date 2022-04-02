package dto

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
	"net/http"
)

func CreateUserDTO(u models.UserSignUp, w http.ResponseWriter) {
	dB := db.ConnectDB()

	_, err := dB.Exec(constants.UserSignUpQuery, u.Email, u.FirstName, u.LastName, u.Age, u.Password)

	if err != nil {
		_, writeErr := w.Write([]byte("Incorrect Values, Try Again"))

		if writeErr != nil {
			log.Println("\n", writeErr)
		}
		w.WriteHeader(http.StatusBadRequest)
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
