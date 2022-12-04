package repository

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func NewDatabaseInstance() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("CONNECTION"))

	if err != nil {
		logrus.Fatal(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatalf("\nPing Error. Connection Lost\n. \n Original Error: %s", err.Error())
	}

	return conn, nil

}
