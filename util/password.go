package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Returns the the bcrypt hash of the password
func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %d", err)
	}
	return string(hashedPassword), nil

}

// To check if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
