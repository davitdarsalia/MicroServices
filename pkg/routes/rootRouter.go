package routes

import (
	"dbPractice/pkg/handlers/auth"
	"dbPractice/pkg/handlers/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func RootRouter() {
	port := os.Getenv("DEF_PORT")
	router := mux.NewRouter()

	router.HandleFunc("/signup", auth.CreateUser).Methods("POST")
	router.HandleFunc("/signin", auth.SignInUser).Methods("POST")

	// Generic Requests
	router.HandleFunc("/allusers", user.GetAllUsers).Methods("GET")
	router.HandleFunc("/allusers/{id}", user.GetUserById).Methods("GET")

	// Authorization Filter
	router.HandleFunc("/userinfo/{id}", user.InfoHandler).Methods("GET")

	startErr := http.ListenAndServe(port, router)

	if startErr != nil {
		log.Fatal(startErr)
	}

}
