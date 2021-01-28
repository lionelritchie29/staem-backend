package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var SecretKey = []byte("secretbgtwoijanganliat")

func GenerateToken(userId int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)

	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["userId"].(float64)
		return int(userId), nil
	} else {
		return -1, err
	}
}