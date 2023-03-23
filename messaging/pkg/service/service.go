package service

import (
	"auth/internal/entities"
	"auth/pkg/repository"
	"github.com/go-playground/validator/v10"
	mq "github.com/rabbitmq/amqp091-go"
)

//go:generate mockery --name=Authorizer
type Authorizer interface {
	CreateUser(u *entities.User) (entities.AuthenticatedUserResponse, error)
	LoginUser(u entities.UserInput) (entities.AuthenticatedUserResponse, error)
	RequestPasswordRecover(u *entities.RecoverPasswordInput) error
	ResetPassword(u *entities.RecoverPasswordInput) error
}

type Service struct {
	Authorizer
}

type AuthService struct {
	repo         repository.AuthDB
	messageQueue *mq.Connection
	validator    *validator.Validate
	credentials  *entities.AWSCredentials
}

func newAuthService(repo repository.Repository, mq *mq.Connection, v *validator.Validate, c *entities.AWSCredentials) *AuthService {
	return &AuthService{repo: repo, messageQueue: mq, validator: v, credentials: c}
}

func New(r repository.Repository, mq *mq.Connection, v *validator.Validate, c *entities.AWSCredentials) *Service {
	return &Service{Authorizer: newAuthService(r, mq, v, c)}
}
