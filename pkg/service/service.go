package service

import (
	todo "github.com/davitdarsalia/BookStoreMicroservices"
	"github.com/davitdarsalia/BookStoreMicroservices/pkg/repository"
)

/* Each interface provides a service, and sort of endpoint handlers */

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
}

type TodoItem interface {
}

// Service - Implements all the services and domain zones in application
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService - Initializes a new service instance. Services need access to DB, so, it accepts repository as an argument
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
