package main

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// verifyToken verify the given token to get its payload.
func VerifyToken(tokenString string) (int64, error) {
	// verify the signature of the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid signature")
		}
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	if err != nil {
		return 0, errors.New("Couldn't handle this token")
	}

	// check the validity of token
	if tokenIsValid := parsedToken.Valid; !tokenIsValid {
		return 0, errors.New("Invalid token")
	}

	// extract claims data in the token
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid token claims")
	}
	userID := int64(claims["userID"].(float64)) // .(float64) means type checking

	return userID, nil
}
