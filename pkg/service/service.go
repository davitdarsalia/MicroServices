package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/go-redis/redis/v8"
)

type Service struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(username, password string) (string, error)
	ResetPassword(r *entities.ResetPassword) (string, error)
	ValidateResetEmail(e *entities.ValidateResetEmail) error
	ResetPasswordProfile(e *entities.ResetPasswordInput) error
	RefreshLogin() int

	// ParseToken - Helper Method For AuthChecker
	ParseToken(token string) (int, error)
}

// Account - Get, Post , Put Or Update
type Account interface {
	GetProfileDetails()
	GetUserInfo()
	GetTrustedDevices()
	GetUserById()

	BlockUser()
	UnblockUser()
	BlockedUsersList()
	UploadProfileImage()
	LogoutSession()

	UpdateProfileDetails()
	UpdateTrustedDevices()

	// SetPasscode - Public/Private Keys
	SetPasscode()
}

type Settings interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewService(repos *repository.Repository, redisConn *redis.Client) *Service {
	return &Service{
		Authorization: NewAuthService(repos, redisConn),
		Account:       NewAccountService(repos, redisConn),
		Transactions:  NewTransactionsService(repos, redisConn),
		Deletions:     NewDeletionsService(repos, redisConn),
	}
}
