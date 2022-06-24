package service

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/davitdarsalia/LendAppBackend/pkg/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"mime/multipart"
)

type Service struct {
	Authorization
	Account
	Settings
}

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(username, password string) (string, error)
	ResetPassword(r *entities.ResetPassword) (string, error)
	ValidateResetEmail(e *entities.ValidateResetEmail) error
	ResetPasswordProfile(e *entities.ResetPasswordInput) error
	RefreshLogin() int

	// ParseToken - Not Endpoint Based Method
	ParseToken(token string) (int, error)
}

// Account TODO - Get, Post , Put Or Update
type Account interface {
	GetProfileDetails() (*entities.ProfileDetails, error)
	GetUserInfo() (*entities.UserInfo, error)
	GetTrustedDevices() ([]entities.TrustedDevices, error)

	BlockUser(b *entities.BlockingUser) error
	UnblockUser(b *entities.UnblockingUser) error
	BlockedUsersList() ([]entities.BlockedUsersList, error)
	UploadProfileImage(c *gin.Context, f multipart.File) error
	LogoutSession() error

	UpdateProfileDetails()
	AddTrustedDevice(d *entities.TrustedDevices) (int, error)

	// SetPasscode - Public/Private Keys
	SetPasscode()
}

// Settings TODO - Settings - Get, Put
type Settings interface {
	GetProfileSettings()
	GetNotificationSettings()
	// GetPaymentOptions - Payments
	GetPaymentOptions()
	GetPrivacySettings()
	GetSecuritySettings()

	UpdateNotificationSettings()
	UpdatePaymentOptions()
	UpdatePrivacySettings()
	UpdateSecuritySettings()
}

type Transactions interface {
}

type Deletions interface {
}

func NewService(repos *repository.Repository, redisConn *redis.Client) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, redisConn),
		Account:       NewAccountService(repos.Account, redisConn),
		Settings:      NewSettingsService(repos.Settings, redisConn),
	}
}
