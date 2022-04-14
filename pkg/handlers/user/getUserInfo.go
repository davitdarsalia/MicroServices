package user

import (
	"dbPractice/pkg/dto/user"
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
	identifier := mux.Vars(r)
	id := identifier["id"]

	secret := os.Getenv("SIGN_KEY")

	token := r.Header.Get("token")

	_, tokenValidation := security.ValidateJWT(token, secret)

	if tokenValidation != nil {
		fmt.Println("Yey")
	}

	if strings.HasPrefix(token, "Bearer") {
		// Trim space
		token = strings.TrimPrefix(token, "Bearer")
	}

	info := user.UserInfoDTO(w, id)

	json, jsonErr := json2.Marshal(info)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	_, writeErr := w.Write(json)

	if writeErr != nil {
		log.Fatal(writeErr)
	}

}
