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
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("Failed To Load Environment Variables: %s", err.Error())
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

	db, err := repository.NewDB(entities.DBConfig{
		// TODO - Move this config into config.yaml
		Host:     "localhost",
		Port:     "5435",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

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
}
