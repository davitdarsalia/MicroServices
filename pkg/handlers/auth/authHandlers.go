package auth

import (
	"dbPractice/pkg/dto/auth"
	"dbPractice/pkg/handlers"
	"dbPractice/pkg/handlers/security"
	"dbPractice/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserSignUp
	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		log.Println(err)
	}
	hashedPassword := security.HashData(userModel.Password)

	userModel.Password = hashedPassword

	createError := auth.CreateUserDTO(userModel, w)

	if createError != false {
		_, writeErr := w.Write([]byte("User Already Exists"))
		if writeErr != nil {
			log.Println("\n", writeErr)
		}
	}
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application-json")
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserSignUp

	err := json.NewDecoder(r.Body).Decode(&userModel)
	if err != nil {
		log.Println(err)
	}

	existence, userId := auth.SignInUserDTO(userModel.Email, userModel.Password)

	if existence == false {
		w.WriteHeader(http.StatusNotFound)
		log.Println("\nOne Of Your Credentials Is Incorrect. Please, Try Again")
		return
	}

	w.WriteHeader(http.StatusOK)
	token := handlers.JwtGenerator(userId)
	byteToken, marshalErr := json.Marshal(token)

	if marshalErr != nil {
		log.Println(marshalErr)
	}

	w.Header().Set("Content-Type", "application-json")
	_, byteWriterError := w.Write(byteToken)
	if byteWriterError != nil {
		log.Println(byteWriterError)
	}
}
