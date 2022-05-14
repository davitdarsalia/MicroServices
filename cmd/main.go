package main

import (
	"github.com/davitdarsalia/LendAppBackend/cache"
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/handler"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/davitdarsalia/LendAppBackend/pkg/service"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	go logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("RootConfig Initialization Error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()

	if err != nil {
		logrus.Fatalf("Error WHile Initializing DataBase Connection; %s", err.Error())
	}

	redisConn := cache.NewRedisCache(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	repos := repository.NewRepository(db)
	services := service.NewService(repos, redisConn)
	handlers := handler.NewHandler(services)

	srv := new(entities.MainServer)

	if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
	}
}

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
