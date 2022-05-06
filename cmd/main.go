package main

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/handler"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("RootConfig Initialization Error: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(entities.MainServer)

	loadEnv()

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	// Add config root directory
	viper.AddConfigPath("configs")
	// Specify name of the root config
	viper.SetConfigName("config")
	// Next, read the config
	return viper.ReadInConfig()
}
