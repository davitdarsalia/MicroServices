package main

import (
	"github.com/davitdarsalia/RestAPI.git/pkg/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	routes.RootRouter()
}

func init() {
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal(envLoadErr)
	}
}
