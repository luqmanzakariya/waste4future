package utils

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the password with pepper from .env
func HashPassword(password string) (string, error) {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		return "", fmt.Errorf("PEPPER is not set in .env")
	}

	// Append pepper to password
	passwordWithPepper := password + pepper

	// Generate bcrypt hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordWithPepper), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with the plain text password.
func CheckPassword(hashedPassword, password string) bool {
	pepper := os.Getenv("PEPPER")
	if pepper == "" {
		log.Println("PEPPER is not set in .env")
		return false
	}

	// Append pepper to password
	passwordWithPepper := password + pepper

	// Compare the hashed password with the provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordWithPepper))
	return err == nil
}
