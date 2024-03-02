/*
 * Проверка верного билета
 */
package handlers

import (
	"fmt"
	"museum/app/contract"
	"museum/app/models/verify"

	"github.com/gofiber/fiber/v2"
)

type verifyTicketRoutes struct{}

// Подключение роутов обработки авторизации
func NewVerifyTicketRoutes() *verifyTicketRoutes {
	return &verifyTicketRoutes{}
}

// CheckTicket godoc
// @Summary      Информация о билете
// @Description  Ссылка зашитая в QR, переход после скана
// @Tags         Проверка билета
// @Accept       json
// @Produce      html
// @Success      200
// @Router       /v1/verify/info/{code} [get]
func (v *verifyTicketRoutes) CheckTicket(ctx *fiber.Ctx) error {
	param := contract.VerifyQrCode{}
	if err := ctx.ParamsParser(&param); err != nil {
		return ctx.Render("verify/info", "")
	}

	ver, err := verify.NewQuery().FindByCode(param.Code)
	if err != nil {
		return ctx.Render("verify/info", "")
	}

	fmt.Println(ver.ID)
	return ctx.Render("verify/info", "")
}
