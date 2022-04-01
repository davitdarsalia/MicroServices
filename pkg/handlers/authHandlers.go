package handlers

import (
	"dbPractice/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var userModel models.UserSignUp

	decoded := json.NewDecoder(r.Body).Decode(&userModel)

	fmt.Println(userModel)

}
