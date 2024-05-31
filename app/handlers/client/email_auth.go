package handlers

import (
	"fmt"
	client_contract "museum/app/contract/client"
	"museum/app/handlers"
	"museum/app/utils"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/gofiber/fiber/v2"
)

type EmailAuthRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEmailAuthRoutes(db *postgres.Postgres, l *logger.Logger) *EmailAuthRoutes {
	return &EmailAuthRoutes{
		db: db,
		l:  l,
	}
}

// Create godoc
// @Summary      Авторизация по почте
// @Description  Отправляем код на почту для подтверждения
// @Tags         Клиент сторона
// @Accept       json
// @Produce      json
// @Success    	 200 {object} client_contract.CreateEmail
// @Failure    	 422 {object} handlers.ErrorStruct
// @Router       /v1/client/sign_in [post].
func (e *EmailAuthRoutes) Create(ctx *fiber.Ctx) error {
	var request client_contract.CreateEmail
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	return ctx.SendStatus(fiber.StatusBadRequest)
}

// Verify godoc
// @Summary      Проверка кода из письма
// @Description  Првеоряем код и авторизиурем
// @Tags         Клиент сторона
// @Accept       json
// @Produce      json
// @Success    	 200 {object} client_contract.VerifyEmail
// @Failure    	 422 {object} handlers.ErrorStruct
// @Router       /v1/client/auth [post].
func (e *EmailAuthRoutes) Verify(ctx *fiber.Ctx) error {
	var request client_contract.VerifyEmail
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	token, err := utils.GenerateToken(2, "email", "user")
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	var twoMonth time.Duration = 1460

	ctx.Cookie(
		&fiber.Cookie{
			Domain:   "museum-ekb.ru",
			Name:     "museum_client_auth",
			Value:    fmt.Sprintf("Bear %s", token),
			SameSite: "Lax",
			Secure:   true,
			HTTPOnly: true,
			Expires:  time.Now().Add(time.Hour * twoMonth).Local(),
		},
	)

	return ctx.SendStatus(fiber.StatusAccepted)
}
