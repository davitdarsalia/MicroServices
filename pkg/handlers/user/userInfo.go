package user

import (
	"dbPractice/pkg/dto"
	"dbPractice/pkg/handlers/security"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	secret := os.Getenv("SIGN_KEY")

	token := r.Header.Get("token")
	fmt.Println(token)

	_, tokenValidation := security.ValidateJWT(token, secret)
	fmt.Println(tokenValidation)

	if strings.HasPrefix(token, "Bearer") {
		// Trim space
		token = strings.TrimPrefix(token, "Bearer")
	}

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		_, writeErr := w.Write([]byte("Not Authorized"))

		if writeErr != nil {
			log.Println(writeErr)
		}
	} else if len(token) > 10 {
		w.WriteHeader(http.StatusOK)
		_, writeErr := w.Write([]byte("UserInfoFetched"))
		if writeErr != nil {
			log.Println(writeErr)
		}
	}

	info := dto.UserInfoDTO()
	fmt.Println(info)

}
