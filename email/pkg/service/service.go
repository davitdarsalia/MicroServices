package service

import (
	"email/pkg/repository"
)

type Mailer interface {
	Dummy()
}

type Service struct {
	Mailer
}

type EmailService struct {
	repo repository.EmailDB
}

func newAuthService(repo repository.Repository) *EmailService {
	return &EmailService{repo: repo}
}

func New(r repository.Repository) *Service {
	return &Service{Mailer: newAuthService(r)}
}
