package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"menuAPI/internal/entities"
	"menuAPI/pkg/handler"
	"menuAPI/pkg/repository"
	"menuAPI/pkg/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "18:08:20",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       true,
	})

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error During Config Initialization: %v", err.Error())
	}

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

	repos := repository.New(db)
	services := service.New(repos)
	handler := handler.New(services)

	s := new(entities.Server)
	if err := s.Run(viper.GetString("port"), handler.DefineRoutes()); err != nil {
		logrus.Fatalf("Error occured while initializing server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("internal/configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
