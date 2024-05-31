package utils

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
)

// Дешифрование данных из url строки.
func Decrypy(value string) (string, error) {
	decode, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher([]byte(secretKey()))
	if err != nil {
		return "", errors.New("12")
	}

	plain := make([]byte, len(decode))
	c.Decrypt(plain, decode)

	return string(plain), nil
}

// Шифрование строки для подставления в url строку.
func Encrypt(value string) (string, error) {
	c, err := aes.NewCipher([]byte(secretKey()))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(value))
	c.Encrypt(out, []byte(value))

	return base64.URLEncoding.EncodeToString(out), nil
}

func secretKey() string {
	return "213"
}
