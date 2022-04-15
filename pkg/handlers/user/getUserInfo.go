package user

import (
	"dbPractice/pkg/dto/user"
	"dbPractice/pkg/handlers"
	"dbPractice/pkg/handlers/security"
	json2 "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	isValid := handlers.TokenIsValid(token)
	if isValid == true {
		secret := os.Getenv("SIGN_KEY")
		identifier := mux.Vars(r)
		id := identifier["id"]
		_, tokenValidation := security.ValidateJWT(token, secret)
		if tokenValidation != nil {
			fmt.Println("Yey")
		}
		if strings.HasPrefix(token, "Bearer") {
			token = strings.TrimPrefix(token, "Bearer")
		}
		info := user.GetUserInfoByIDDTO(w, id)
		json, jsonErr := json2.Marshal(info)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		_, writeErr := w.Write(json)
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	} else if isValid == false {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
