package repository

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(username, password string) (entities.User, error)
	ResetPassword(p *entities.ResetPassword) (string, error)
	ValidateResetEmail(p *entities.ValidateResetEmail) error
	ResetPasswordProfile(p *entities.ResetPasswordInput) error
}

// Account - TODO // Implement These Methods
type Account interface {
	GetProfileDetails(userID *int) (*entities.ProfileDetails, error)
	GetUserInfo(userID *int) (*entities.UserInfo, error)
	GetTrustedDevices(userID *int) ([]entities.TrustedDevices, error)
	GetImages(userID *int) ([]entities.Image, error)

	BlockUser(userID *int, u *entities.BlockingUser) error
	UnblockUser(userID *int, u *entities.UnblockingUser) error
	BlockedUsersList(userID *int) ([]entities.BlockedUsersList, error)
	UploadProfileImage(f string, userID int, uploadTime *string) error
	LogoutSession()

	UpdateProfileDetails()
	AddTrustedDevice(userID *int, t *entities.TrustedDevices) (int, error)

	// SetPasscode - Public/Private Keys
	SetPasscode()

	WriteAccountIpToDB(userID string) error
}

// Settings - TODO // Implement These Methods
type Settings interface {
	GetProfileSettings()
	GetNotificationSettings()
	// GetPaymentOptions - Payments
	GetPaymentOptions()
	GetPrivacySettings()
	GetSecuritySettings()

	UpdateNotificationSettings(userID *int, s *entities.NotificationSettings) error
	UpdatePaymentSettings(userID *int, s *entities.PaymentSettings) error
	UpdateSecuritySettings(userID *int, s *entities.SecuritySettings) error

	// UpdatePrivacySettings - TODO - Implement DB And Method
	UpdatePrivacySettings() error

	WriteSettingsIpToDB(userID string) error
}

type Transactions interface {
}

type Deletions interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountPostgres(db),
		Settings:      NewSettingsPostgres(db),
	}
}
