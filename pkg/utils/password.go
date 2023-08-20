package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, pepper string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password+pepper), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	return hash, nil
}

func VerifyPassword(hashedPassword []byte, password string, pepper string) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, []byte(password+pepper))
}
