package repository

type Repository struct {
	Authorization
	Account
	Transactions
	Deletions
}

type Authorization interface {
}

type Account interface {
}

type Transactions interface {
}

type Deletions interface {
}

func NewRepository() *Repository {
	return &Repository{}
}
