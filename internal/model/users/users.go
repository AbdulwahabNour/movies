package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"create_at"`
	Name     string    `json:"name" validate:"required,max=100"`
	Email    string    `json:"email" validate:"required,email,max=100"`
	Password password  `json:"password" validate:"required"`
	Actived  bool      `json:"activated"`
	Version  string    `json:"version"`
}
type password struct {
	plainText *string
	hashed    []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.plainText = &text
	p.hashed = hash

	return nil
}
func (p *password) ComparePasswords(text string) (bool, error) {

	err := bcrypt.CompareHashAndPassword(p.hashed, []byte(text))
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
