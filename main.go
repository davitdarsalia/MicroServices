package main

import (
	"dbPractice/pkg/routes"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	routes.RootRouter()
}

func init() {
	envError := godotenv.Load()

	if envError != nil {
		log.Fatal(envError)
	}
}
