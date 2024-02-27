/*
 * Промежуточные проверки перед обратокой контроллеров
 */
package middleware

import (
	"fmt"
	"museum/app/handlers/helpers"
	users "museum/app/models/user"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type autorizeHeader struct {
	Token string `reqHeader:"Authorization"`
}

// Проверка аунтификации по jwt токену
func AuthAccess(ctx *fiber.Ctx) error {
	headers := new(autorizeHeader)
	if err := ctx.ReqHeaderParser(headers); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token, ok := existsToken(clearToken(headers.Token))
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	fmt.Println(token.Claims.(jwt.MapClaims))
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	user, ok := existsUser(id)
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	ctx.Locals("currentUser", user)

	return ctx.Next()
}

// Парсинг токена
func existsToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.JwtSecretKey()), nil
	})
	if err != nil || !token.Valid {
		return token, false
	}

	return token, true
}

// Очистка лишнего из строки с токеном
func clearToken(authField string) string {
	splited := strings.Split(authField, helpers.JwtSeparateKey())
	token := splited[len(splited)-1]
	return strings.TrimSpace(token)
}

// Поиск сщуестующего пользователя
func existsUser(id int) (*users.User, bool) {
	user, err := users.NewQuery().FindByID(id)
	if err != nil {
		return nil, false
	}
	return user, true
}
