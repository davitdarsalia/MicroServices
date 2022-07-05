package entities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type User struct {
	UserID         int    `json:"user_id"`
	PersonalNumber string `json:"personal_number" binding:"required"`
	PhoneNumber    string `json:"phonenum" binding:"required"`
	UserName       string `json:"user_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	IpAddress      string `json:"ip_address" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Salt           []byte `json:"salt"`
}

type UserRegInput struct {
	PersonalNumber string `json:"personal_number" binding:"required"`
	PhoneNumber    string `json:"phonenum" binding:"required"`
	UserName       string `json:"user_name" binding:"required"`
	Email          string `json:"email" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	IpAddress      string `json:"ip_address" binding:"required"`
	Password       string `json:"password" binding:"required"`
}

type UserInput struct {
	UserId   string `json:"user_id" db:"user_id"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserInputWithoutID struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CustomToken struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Ip       string `json:"ip_address"`
}

type ResetPassword struct {
	Email          string `json:"email" db:"email" binding:"required"`
	PersonalNumber string `json:"personal_number" db:"personal_number" binding:"required"`
	UserName       string `json:"username" db:"username" binding:"required"`
}

type ValidateResetEmail struct {
	PersonalNumber string `json:"personal_number" db:"personal_number" binding:"required"`
	ValidationCode string `json:"validation_code" binding:"required"`
	NewPassword    string `json:"password" db:"password" binding:"required"`
}

type ResetPasswordInput struct {
	UserName    string `json:"username" db:"username" binding:"required"`
	NewPassword string `json:"new_password" db:"password" binding:"required"`
}

type RefreshLoginInput struct {
	RefreshToken string `json:"refresh_token"`
}

var (
	Header  = "Authorization"
	SignKey = os.Getenv("SIGN_IN_KEY")
	UserCtx = os.Getenv("USER_CONTEXT")

	SendMailFrom = "d.app.geo1@gmail.com"
	MailHost     = "smtp.gmail.com"
	MailPort     = "587"

	MailAuthPassword = "David.1998!"

	OtpExpireDate = time.Minute * 1
)

var (
	MailAddress = fmt.Sprintf("%s : %s", MailHost, MailPort)
)
