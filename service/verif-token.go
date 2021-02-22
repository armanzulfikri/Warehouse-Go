package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

// VerifyToken func
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	secret := os.Getenv("SECRET_TOKEN")
	// when authorization is nil
	if tokenString == "" {
		return nil, fmt.Errorf("Missing Authorization Header")
	}

	// decode token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return false, fmt.Errorf("There was an error")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	// when token is invalid
	if !token.Valid {
		return nil, fmt.Errorf("token invalid")
	}

	return token.Claims.(jwt.MapClaims), nil
}
