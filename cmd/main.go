package main

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	srv := new(entities.MainServer)
	loadEnv()

	srv.Run(os.Getenv("PORT"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
