package dto

import (
	"dbPractice/pkg/constants"
	"dbPractice/pkg/db"
	"dbPractice/pkg/models"
	"log"
)

func CreateUserDTO(u models.UserSignUp) {
	dB := db.ConnectDB()

	_, err := dB.Exec(constants.UserSignUpQuery, u.Email, u.FirstName, u.LastName, u.Age, u.SecretWord)

	if err != nil {
		log.Fatal(err)
	}
}
