package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(adminId uint, name string) (string, error) {
	// create the claims
	claims := jwt.MapClaims{}
	claims["admin_id"] = adminId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
