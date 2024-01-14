package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
)

func JwtAuthentication(accesToken string) bool {

	return true
}

func ParseAccessToken(accessToken string) *jwt.Token {
	parsedAccessToken, _ := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY_JWT")), nil
	})

	return parsedAccessToken
}
