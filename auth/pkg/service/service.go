package service

import (
	"auth/internal/entities"
	"auth/pkg/repository"
	"github.com/go-playground/validator/v10"
	mq "github.com/rabbitmq/amqp091-go"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorizer interface {
	CreateUser(u entities.User) (entities.AuthenticatedUserResponse, error)
	LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error)
	RecoverPassword(u entities.RecoverPasswordInput) error

	// CheckToken -  Helper method
	CheckToken(authToken, signKey string) (string, error)
}

type Service struct {
	Authorizer
}

type AuthService struct {
	repo         repository.AuthDB
	messageQueue *mq.Connection
	validator    *validator.Validate
}

func newAuthService(repo repository.Repository, mq *mq.Connection, v *validator.Validate) *AuthService {
	return &AuthService{repo: repo, messageQueue: mq, validator: v}
}

func New(r repository.Repository, mq *mq.Connection, v *validator.Validate) *Service {
	return &Service{Authorizer: newAuthService(r, mq, v)}
}
