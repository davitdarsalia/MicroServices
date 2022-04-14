package routes

import (
	"dbPractice/pkg/handlers/auth"
	"dbPractice/pkg/handlers/user"
	"github.com/gorilla/handlers"
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
	router.HandleFunc("/userinfo/{id}", user.GetUserInfo).Methods("GET")
	router.HandleFunc("/increase_rating/{id}", user.IncreaseRating).Methods("POST")
	router.HandleFunc("/increase_balance/{id}", user.IncreaseBalance).Methods("POST")

	// Payment - Half Anonymous Transaction
	router.HandleFunc("/transaction", user.MakeTransaction).Methods("POST")

	router.Use()

	// Cors Policy Resolver
	startErr := http.ListenAndServe(port, handlers.CORS()(router))

	if startErr != nil {
		log.Fatal(startErr)
	}

}
