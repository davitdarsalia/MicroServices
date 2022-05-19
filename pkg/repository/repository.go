package repository

import (
	"github.com/davitdarsalia/LendAppBackend/entities"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
	RegisterUser(u *entities.User) (int, error)
	CheckUser(username, password string) (entities.User, error)
	ResetPassword(p *entities.ResetPassword) (string, error)
	ValidateResetEmail(p *entities.ValidateResetEmail) error
	ResetPasswordProfile(p *entities.ResetPasswordInput) error
	RefreshLogin()
}

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

type Transactions interface {
}

type Deletions interface {
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Account:       NewAccountPostgres(db),
		Transactions:  NewTransactionsPostgres(db),
		Deletions:     NewDeletionsPostgres(db),
	}
}
