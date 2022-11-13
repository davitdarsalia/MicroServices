package main

import (
	"os"

	"github.com/davitdarsalia/auth/internal"
	"github.com/davitdarsalia/auth/pkg/handler"
	"github.com/davitdarsalia/auth/pkg/repository"
	"github.com/davitdarsalia/auth/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var globalFormatter = new(logrus.JSONFormatter)

// @title Authentication Server
// @version 0.0.1
// @description Endpoints For Authorization, Authentication

// @host: localhost: 8100
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus.SetFormatter(globalFormatter)

	defer func() {

	}()

	/* Instances Been Initializing Here */
	db, err := repository.NewDatabaseInstance()
	if err != nil {
		logrus.Fatalf("Error While Initializing DataBase Connection; %s", err.Error())
	}

	repos := repository.New(db)
	services := service.New(repos)
	handlers := handler.New(services)

	// Server Runs In Separate GoRoutine

	go func() {
		if err := new(internal.AuthServer).Run(os.Getenv("PORT"), handlers.Routes()); err != nil {

		}
	}()

}

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed To Initialize Enviroment Variables: %s", err)
	}
}
