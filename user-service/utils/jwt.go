package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// UserInfo holds the user information extracted from the JWT
type UserInfo struct {
	ID        int
	Name      string
	Email     string
	AddressID string
}

func GenerateJWT(id int, name string, email string, addressID string) (string, error) {
	claims := jwt.MapClaims{
		"id":         id,
		"name":       name,
		"email":      email,
		"address_id": addressID,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("error secret is empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*UserInfo, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT_SECRET environment variable is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok := claims["id"].(float64)
		if !ok {
			return nil, errors.New("invalid id in token")
		}

		name, ok := claims["name"].(string)
		if !ok {
			return nil, errors.New("invalid name in token")
		}

		email, ok := claims["email"].(string)
		if !ok {
			return nil, errors.New("invalid email in token")
		}

		addressID, ok := claims["address_id"].(string)
		if !ok {
			return nil, errors.New("invalid address_id in token")
		}

		userInfo := &UserInfo{
			ID:        int(id),
			Name:      name,
			Email:     email,
			AddressID: addressID,
		}

		return userInfo, nil
	}

	return nil, errors.New("invalid token")
}
