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
	// DataBase Commands
	dto.CreateUserDTO(userModel)

	defer func() {
		_, err = w.Write([]byte("User Created"))
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application-json")

	}()
	w.WriteHeader(http.StatusCreated)
}
