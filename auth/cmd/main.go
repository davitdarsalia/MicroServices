package main

import (
	"auth/internal/entities"
	"auth/internal/outerServices"
	"auth/pkg/handler"
	"auth/pkg/repository"
	"auth/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	. "log"
	"os"
)

// @title Authentication Server
// @version 1.0
// @description Authentication Server documentation. Includes Authentication | Authorization, Reset Password, Messaging Queues
// @host: localhost:8080
// @BasePath localhost:8080/authServer
// @securityDefinitions.apikey ApiKetAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	logFile, err := os.Create("info.log")
	logrus.SetOutput(logFile)
	defer logFile.Close()

	if err != nil {
		Println("Unable to create info.log: \n", err.Error())
	}

	if err != nil {
		Printf("Failed To Load Environment Variables: %s", err.Error())
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "18:08:20",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       true,
	})

	db, err := repository.NewDB()

	if err != nil {
		logrus.Fatalf("Failed to create database: %v", err.Error())
	}

	v := validator.New()
	mq := outerServices.MqConnection()
	repos := repository.New(db)
	services := service.New(repos, mq, v)
	h := handler.New(services)
	s := new(entities.Server)

	if err := s.Run(os.Getenv("AUTH_SERVER_PORT"), h.DefineRoutes()); err != nil {
		logrus.Fatalf("Error occured while initializing server: %s", err.Error())
	}
	// TODO - Close rabbitMQ connection on graceful shutdown
}
