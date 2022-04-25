package repository

import (
	todo "github.com/davitdarsalia/BookStoreMicroservices"
	"github.com/jmoiron/sqlx"
)

/* Each interface provides a repository service, to communicate with DataBase */

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
}

type TodoItem interface {
}

// Repository - Implements all the services and domain zones in application
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository - Initializes a new repository instance
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      nil,
	}
}
