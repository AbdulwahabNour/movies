package users

import (
	"errors"
	"time"

	"github.com/AbdulwahabNour/movies/internal/model/token"
)

var AnonymousUser = &User{}

type User struct {
	ID             int64     `json:"id"`
	CreateAt       time.Time `json:"create_at" `
	Name           string    `json:"name" validate:"required,max=100" `
	Email          string    `json:"email" validate:"required,email,max=100"`
	Password       string    `json:"password,omitempty" validate:"required,min=8,max=50" `
	HashedPassword []byte    `json:"-"  validate:"required"`
	Activated      *bool     `json:"activated,omitempty"`
	Version        string    `json:"version"`
}
type SignUpInput struct {
	Name            string `json:"name" validate:"required,max=100"`
	Email           string `json:"email" validate:"required,email,max=100"`
	Password        string `json:"password" validate:"required,min=8,max=50" `
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=50" `
}
type SignIn struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=8,max=200" `
}

type UserWithToken struct {
	User  *User            `json:"user"`
	Token *token.TokenPair `json:"token"`
}

func (u *User) SanitizePassword() {
	u.Password = ""
}
func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (u *SignUpInput) Check() error {
	if u.Password != u.ConfirmPassword {
		return errors.New("passwords do not match")
	}

	return nil
}
func (u *SignUpInput) Map() User {

	return User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
