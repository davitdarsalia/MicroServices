package main

import (
	"auth/internal/entities"
	"auth/internal/outerServices"
	"auth/pkg/handler"
	"auth/pkg/repository"
	"auth/pkg/service"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Authentication Server
// @version 1.0
// @description Authentication Server documentation.
// Includes Authentication | Authorization, Reset Password, Messaging Queues
// @host: localhost:8080
// @BasePath localhost:8080/authServer
// @securityDefinitions.apikey ApiKetAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("Failed to load environment variables: %s", err.Error())
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

	secrets, err := outerServices.GetSecret()
	if err != nil {
		logrus.Fatalf("Failed to create database: %v", err.Error())
	}

	conn, err := pgxpool.Connect(context.Background(), secrets.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	v := validator.New()
	mq := outerServices.MqConnection(&secrets.RabbitMQConn)
	repos := repository.New(conn, &secrets)
	services := service.New(repos, mq, v, &secrets)
	h := handler.New(services)
	s := new(entities.Server)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.Run(secrets.ServerPort, h.DefineRoutes()); err != nil {
			logrus.Fatalf("Error occurred while initializing server: %s", err.Error())
		}
	}()

	<-signalChan

	mq.Close()
	conn.Close()
}
