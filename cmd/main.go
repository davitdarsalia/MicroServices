package main

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/handler"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	srv := new(entities.MainServer)

	rootHandler := new(handler.Handler)
	loadEnv()

	if err := srv.Run(os.Getenv("PORT"), rootHandler.InitRoutes()); err != nil {
		log.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
