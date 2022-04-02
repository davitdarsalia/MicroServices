package handlers

import (
	"dbPractice/pkg/dto"
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

	createError := dto.CreateUserDTO(userModel, w)

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

	existence := dto.SignInUserDTO(userModel.Email, userModel.Password)
	if existence == false {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("\nOne Of Your Credentials Is Incorrect. Please, Try Again \n")
		return
	}

	w.WriteHeader(http.StatusOK)
	token := JwtGenerator()

	byteResponse, marshalErr := json.Marshal([]byte(token))
	if marshalErr != nil {
		log.Println(marshalErr)
	}

	w.Header().Set("Content-Type", "application-json")
	_, byteWriterError := w.Write(byteResponse)
	if byteWriterError != nil {
		log.Println(byteWriterError)
	}
}
