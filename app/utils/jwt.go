package utils

import (
	"museum/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Генерация токена авторизации.
func GenerateToken(id int, email string, role string) (string, error) {
	var twoDays time.Duration = 72
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * twoDays).Unix(),
		"role":  role,
	}).SignedString([]byte(JwtSecretKey()))

	if err != nil {
		return "", err
	}

	return token, nil
}

func JwtSecretKey() string {
	cfg, _ := config.GetConf()

	return cfg.HTTP.JwtSecretKey
}

func JwtSeparateKey() string {
	cfg, _ := config.GetConf()

	return cfg.HTTP.JwtSeparateKey
}
