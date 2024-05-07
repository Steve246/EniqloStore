package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PasswordRepository interface {
	HashAndSavePassword(userPass string) (string, error)
	VerifyPassword(dbPass []byte, userPass []byte) error
}

type passwordRepository struct {
	db *gorm.DB
}

func (u *passwordRepository) HashAndSavePassword(userPass string) (string, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (u *passwordRepository) VerifyPassword(dbPass []byte, userPass []byte) error {

	result := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if result != nil {
		// Passwords do not match
		return result
	}

	return nil // Passwords match; user authenticated
}

func NewPasswordRepository(db *gorm.DB) PasswordRepository {
	repo := new(passwordRepository)
	repo.db = db
	return repo
}
