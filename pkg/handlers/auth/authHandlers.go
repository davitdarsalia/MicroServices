package auth

import (
	"dbPractice/pkg/dto/auth"
	"dbPractice/pkg/handlers"
	"dbPractice/pkg/handlers/security"
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
	hashedPassword := security.HashData(userModel.Password)
	userModel.Password = hashedPassword
	createError := auth.CreateUserDTO(userModel, w)
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
	existence, userId := auth.SignInUserDTO(userModel.Email, userModel.Password)
	if existence == false {
		w.WriteHeader(http.StatusNotFound)
		log.Println("\nOne Of Your Credentials Is Incorrect. Please, Try Again")
		return
	}
	w.WriteHeader(http.StatusOK)
	token := handlers.JwtGenerator(userId)
	refreshToken, refreshTokenErr := handlers.RefreshToken()
	if refreshTokenErr != nil {
		log.Fatal(refreshTokenErr)
	}
	token.RefreshToken = refreshToken
	byteToken, marshalErr := json.Marshal(token)
	if marshalErr != nil {
		log.Println(marshalErr)
	}
	w.Header().Set("Content-Type", "application-json")
	_, byteWriterError := w.Write(byteToken)
	if byteWriterError != nil {
		log.Println(byteWriterError)
	}
}

func LogOut(w http.ResponseWriter, r *http.Request) {
	var token string
	decodeErr := json.NewDecoder(r.Body).Decode(&token)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			log.Fatal(closeErr)
		}
	}()
}

func RefreshLogin(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	isValid := handlers.TokenIsValid(token)
	if isValid == true {
		w.WriteHeader(http.StatusOK)
		newToken, _ := handlers.RefreshToken()
		_, writeErr := w.Write([]byte(newToken))
		if writeErr != nil {
			log.Fatal(writeErr)
		}
	} else if isValid == false {
		w.WriteHeader(http.StatusForbidden)
		_, writeErr := w.Write([]byte("Access Forbidden"))
		if writeErr != nil {
			log.Println(writeErr)
		}
	}
}
