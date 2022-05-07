package main

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/handler"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("RootConfig Initialization Error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()

	if err != nil {
		logrus.Fatalf("Error WHile Initializing DataBase Connection; %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(entities.MainServer)

	loadEnv()

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
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
