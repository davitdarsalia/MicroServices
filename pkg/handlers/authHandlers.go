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
		log.Fatal(err)
	}

	dto.CreateUserDTO(userModel, w)

	defer func() {
		_, err = w.Write([]byte("User Created"))
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application-json")

	}()
	w.WriteHeader(http.StatusCreated)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserSignUp
	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		log.Fatal(err)
	}

	existence := dto.SignInUserDTO(userModel.Email, userModel.Password)

	if existence {
		w.WriteHeader(http.StatusOK)
		log.Println("Correct Credentials")
	} else if existence == false {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("\nOne Of Your Credentials Is Incorrect. Please, Try Again \n")
	}

}
