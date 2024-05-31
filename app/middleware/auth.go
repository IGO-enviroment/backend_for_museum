package middleware

import (
	"museum/app/models"
	"museum/app/utils"
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

// Инициализация структуры для проверка аунтификации.
func NewAuthAccess(db *postgres.Postgres, l *logger.Logger) *AuthAccess {
	return &AuthAccess{
		db: db,
		l:  l,
	}
}

// Проверка аунтификации по jwt токену.
func (a *AuthAccess) Аuthorized(ctx *fiber.Ctx) error {
	headers := new(autorizeHeader)
	if err := ctx.ReqHeaderParser(headers); err != nil {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	strToken, ok := a.clearToken(headers.Token)
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	token, ok := a.existsToken(strToken)
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

// Проверка на наличие роли админа.
func (a *AuthAccess) AdminAccess(ctx *fiber.Ctx) error {
	var ok bool
	_, ok = ctx.Locals("currentUser").(*models.User)

	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	ok = false
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	return ctx.Next()
}

// Проверка на наличие роли супер админа.
func (a *AuthAccess) SuperAdminAccess(ctx *fiber.Ctx) error {
	var ok bool
	_, ok = ctx.Locals("currentUser").(*models.User)

	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	ok = false
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	return ctx.Next()
}

// Парсинг токена.
func (a *AuthAccess) existsToken(tokenString string) (*jwt.Token, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(utils.JwtSecretKey()), nil
	})

	if err != nil || !token.Valid {
		return token, false
	}

	return token, true
}

// Очистка лишнего из строки с токеном.
func (a *AuthAccess) clearToken(authField string) (string, bool) {
	bearAndToken := 2
	splited := strings.Split(authField, utils.JwtSeparateKey())

	if len(splited) != bearAndToken {
		return "", false
	}

	token := splited[len(splited)-1]

	return strings.TrimSpace(token), true
}

// Поиск сщуестующего пользователя.
func (a *AuthAccess) existsUser(_ int) (*models.User, bool) {
	var err error
	if err != nil {
		return nil, false
	}

	return nil, true
}
