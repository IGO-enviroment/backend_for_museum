package utils

import (
	"museum/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Генерация токена авторизации
func GenerateToken(id, email string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}).SignedString([]byte(JwtSecretKey()))
	if err != nil {
		return "", err
	}
	return token, nil
}

func JwtSecretKey() string {
	return config.GetConf().HTTP.JwtSecretKey
}

func JwtSeparateKey() string {
	return config.GetConf().HTTP.JwtSeparateKey
}
