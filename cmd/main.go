package main

import (
	"context"
	"fmt"
	todo "github.com/davitdarsalia/BookStoreMicroservices"
	"github.com/davitdarsalia/BookStoreMicroservices/pkg/handlers"
	"github.com/davitdarsalia/BookStoreMicroservices/pkg/repository"
	"github.com/davitdarsalia/BookStoreMicroservices/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList App

// @host localhost:8080
// @basePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error Of Initialization Configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed To Load Enviroment Variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		UserName: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Failed To Initialize DB. %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	s := new(todo.Server)
	go func() {
		if err := s.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running Server on port %s", "8080")
		}
	}()
	logrus.Println("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("App Shutting Down")

	s.ShutDown(context.Background())

	if err := db.Close(); err != nil {
		fmt.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
