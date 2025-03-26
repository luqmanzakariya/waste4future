package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id string) (string, error) {
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 1).Unix(),
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

func ValidateJWT(tokenString string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", errors.New("JWT_SECRET environment variable is not set")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the "id" claim exists and is a string
		id, ok := claims["id"].(string)
		if !ok {
			// If the "id" is not a string, try converting it from a float64 (common in JWTs)
			idFloat, ok := claims["id"].(float64)
			if !ok {
				return "", errors.New("invalid id in token: id is not a string or number")
			}
			// Convert the float64 to a string
			id = strconv.FormatFloat(idFloat, 'f', -1, 64)
		}
		return id, nil
	}

	return "", errors.New("invalid token") // Generic invalid token error
}
