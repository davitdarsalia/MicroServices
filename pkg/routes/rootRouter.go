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

	router.HandleFunc("/transactions", user.TransactionsHandler).Methods("GET")
	router.HandleFunc("/transactions/{:id}", user.TransactionsHandler).Methods("GET")

	startErr := http.ListenAndServe(port, router)

	if startErr != nil {
		log.Fatal(startErr)
	}

}
