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
	defer func(conn *pgx.Conn, ctx context.Context) {
		closeErr := conn.Close(ctx)
		if err != nil {
			logrus.Fatal(closeErr)
		}
	}(conn, context.Background())

	if err != nil {
		logrus.Fatal(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatalf("\nPing Error. Connection Lost\n. \n Original Error: %s", err.Error())
	}

	return conn, nil

}
