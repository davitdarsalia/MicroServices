package user

import (
	"dbPractice/pkg/dto/user"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	identifier := mux.Vars(r)
	id := identifier["id"]
	u := user.UserByIdDTO(w, id)
	userJson, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	_, writeErr := w.Write(userJson)
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
