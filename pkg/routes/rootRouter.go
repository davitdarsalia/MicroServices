package routes

import (
	"dbPractice/pkg/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RootRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.CreateUser).Methods("POST")

	// Router presented as a handler
	startErr := http.ListenAndServe(":8080", router)

	if startErr != nil {
		log.Fatal(startErr)
	}

}
