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
	RefreshLogin()
}

// Account - TODO // Implement These Methods
type Account interface {
	GetProfileDetails(userID *int) (*entities.ProfileDetails, error)
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

// Settings - TODO // Implement These Methods
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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountPostgres(db),
		Settings:      NewSettingsPostgres(db),
	}
}
