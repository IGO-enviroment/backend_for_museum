/*
 * Промежуточные проверки перед обратокой контроллеров
 */
package middleware

import (
	"museum/app/handlers/helpers"
	users "museum/app/models/user"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthAccess struct {
	db *postgres.Postgres
	l  *logger.Logger
}

type autorizeHeader struct {
	Token string `reqHeader:"Authorization"`
}

func NewAuthAccess(db *postgres.Postgres, l *logger.Logger) *AuthAccess {
	return &AuthAccess{
		db: db,
		l:  l,
	}
}

// Проверка аунтификации по jwt токену
func (a *AuthAccess) Аuthorized(ctx *fiber.Ctx) error {
	headers := new(autorizeHeader)
	if err := ctx.ReqHeaderParser(headers); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	token, ok := a.existsToken(a.clearToken(headers.Token))
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	id, err := strconv.Atoi(claims["id"].(string))
	if err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	user, ok := a.existsUser(id)
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	ctx.Locals("currentUser", user)

	return ctx.Next()
}

// Проверка на наличие роли админа
func (a *AuthAccess) AdminAccess(ctx *fiber.Ctx) error {
	var ok bool
	user, ok := ctx.Locals("currentUser").(*users.User)

	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	ok = user.Query.IsAdmin()
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	return ctx.Next()
}

// Парсинг токена
func (a *AuthAccess) existsToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(helpers.JwtSecretKey()), nil
	})
	if err != nil || !token.Valid {
		return token, false
	}

	return token, true
}

// Очистка лишнего из строки с токеном
func (a *AuthAccess) clearToken(authField string) string {
	splited := strings.Split(authField, helpers.JwtSeparateKey())
	token := splited[len(splited)-1]

	return strings.TrimSpace(token)
}

// Поиск сщуестующего пользователя
func (a *AuthAccess) existsUser(id int) (*users.User, bool) {
	user, err := users.NewQuery().FindByID(id)
	if err != nil {
		return nil, false
	}

	return user, true
}
