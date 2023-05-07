package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dvnggak/miniProject/constants"
)

func CreateTokenAdmin(adminId uint, name string) (string, error) {
	// create the claims
	claims := jwt.MapClaims{}
	claims["admin_id"] = adminId
	claims["name"] = name
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(constants.SECRET_JWT))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CreateTokenUser(userId uint, name string, ID_Number string) (string, error) {
	// create the claims
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["id_number"] = ID_Number
	claims["name"] = name
	claims["role"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(constants.SECRET_JWT))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
