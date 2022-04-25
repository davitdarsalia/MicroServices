package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	"log"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.UserName, c.Password, c.DBName, c.SSLMode))

	if err != nil {
		log.Fatalf("Failed To Connect DB. %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping Error. Connection Lost. %s", err.Error())
	}

	return db, nil
}
