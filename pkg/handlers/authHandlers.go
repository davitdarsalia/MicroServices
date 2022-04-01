package handlers

import (
	"dbPractice/pkg/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserSignUp

	err := json.NewDecoder(r.Body).Decode(&userModel)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userModel)

}
