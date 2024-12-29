package utils

import (
	"fmt"
	"os"
	"story-plateform/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwt_secrete = []byte(os.Getenv("JWT_SECRETE_KEY"))

func GenerateJwtToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(jwt_secrete)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwt_secrete, nil
	})

	fmt.Println("ParseToken - token -", token)

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
