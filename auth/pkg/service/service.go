package service

import (
	"auth/internal/entities"
	"auth/pkg/repository"
	mq "github.com/rabbitmq/amqp091-go"
)

type Authorizer interface {
	CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error)
	LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error)
	RecoverPassword(u entities.RecoverPasswordInput) error
}

type Service struct {
	Authorizer
}

type AuthService struct {
	repo         repository.AuthDB
	messageQueue *mq.Connection
}

func newAuthService(repo repository.Repository, mq *mq.Connection) *AuthService {
	return &AuthService{repo: repo, messageQueue: mq}
}

func New(r repository.Repository, mq *mq.Connection) *Service {
	return &Service{Authorizer: newAuthService(r, mq)}
}
