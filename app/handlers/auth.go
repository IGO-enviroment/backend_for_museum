/*
 * Обработка запросов авторизации
 */
package handlers

import (
	"fmt"

	entity "museum/app/entity/auth"
	models "museum/app/models/user"
	"museum/app/utils"
	"museum/pkg/postgres"

	"museum/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

// Подключение роутов обработки авторизации
func NewAuthRoutes(db *postgres.Postgres, l *logger.Logger) *authRoutes {
	return &authRoutes{
		db: db,
		l:  l,
	}
}

// SignUp godoc
// @Summary      Регистрация
// @Description  Ввод почты, пароля и подтверждение пароля
// @Tags         Авторизация
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200
// @Router       /v1/auth/sign_up [post]
func (r *authRoutes) SignUp(ctx *fiber.Ctx) error {
	var req entity.SignUpEntity
	if err := ctx.BodyParser(&req); err != nil {
		return fmt.Errorf("body parser: %w", err)
	}

	err := validator.New().Struct(req)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	token, err := utils.GenerateToken("1", "email")
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"token": token})
}

// SignIn godoc
// @Summary      Вход
// @Description  Проверка почты и пароля, генерация токена
// @Tags         Авторизация
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/auth/sign_in [post]
func (r *authRoutes) SignIn(ctx *fiber.Ctx) error {
	// user := models.NewUser()
	token, err := utils.GenerateToken("1", "email")
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{"token": token})
}

// GetMe godoc
// @Summary      Получение текущего пользователя
// @Description  Поиск пользователя по токену
// @Tags         Авторизация
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/auth/me [get]
func (r *authRoutes) GetMe(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("currentUser").(*models.User)
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	return ctx.Status(fiber.StatusGone).JSON(fiber.Map{"email": user.Email})
}
