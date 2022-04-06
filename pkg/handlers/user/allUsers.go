package user

import (
	"dbPractice/pkg/dto"
	"encoding/json"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := dto.AllUsersDTO(w)

	if err != nil {
		log.Println(err)
	}

	allUsers, err := json.Marshal(users)

	_, writeError := w.Write(allUsers)

	if writeError != nil {
		log.Fatal(writeError)
	}

}
