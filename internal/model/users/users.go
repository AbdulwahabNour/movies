package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int64     `json:"id"`
	CreateAt       time.Time `json:"create_at"`
	Name           string    `json:"name" validate:"required,max=100"`
	Email          string    `json:"email" validate:"required,email,max=100"`
	Password       string    `json:"password" validate:"required,min=10,max=200" `
	HashedPassword []byte    `json:"-"`
	Actived        bool      `json:"activated"`
	Version        string    `json:"version"`
}

func (u *User) SetHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashedPassword = hash
	return nil
}
func (u *User) ComparePasswords(text string) (bool, error) {

	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(text))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
