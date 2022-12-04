package main

import (
	"context"
	"github.com/davitdarsalia/payment/internal"
	"github.com/davitdarsalia/payment/internal/cache"
	"github.com/davitdarsalia/payment/pkg/handler"
	"github.com/davitdarsalia/payment/pkg/repository"
	"github.com/davitdarsalia/payment/pkg/service"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var globalFormatter = new(logrus.JSONFormatter)

func main() {
	logrus.SetFormatter(globalFormatter)

	/* Instances Been Initializing Here */
	db, err := repository.NewDatabaseInstance()
	if err != nil {
		logrus.Fatalf("Error While Initializing DataBase Connection; %s", err.Error())
	}

	redisCache := cache.New(&redis.Options{
		Addr: os.Getenv("REDIS_PORT"),
		DB:   0,
	})
	repos := repository.New(db)
	services := service.New(repos, redisCache)
	handlers := handler.New(services)

	// Server Runs In Separate GoRoutine

	s := new(internal.AuthServer)

	go func() {
		if err := s.Run(os.Getenv("PORT"), handlers.Routes()); err != nil {
			logrus.Fatalf("Error While Running Server On Port %s", os.Getenv("PORT"))
		}
	}()

	quitSignal := make(chan os.Signal, 1)

	signal.Notify(quitSignal, syscall.SIGTERM, syscall.SIGINT)

	<-quitSignal

	/* Graceful Shutdown */

	if err := s.Kill(context.Background()); err != nil {
		logrus.Errorf("Failed To Shut Down Server: \n %s", err.Error())
	}

	if err := db.Close(context.Background()); err != nil {
		logrus.Errorf("Failed To Close DB: \n %s", err.Error())
	}

}

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed To Initialize Enviroment Variables: %s", err)
	}

}