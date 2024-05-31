package handlers

import (
	"github.com/gofiber/fiber/v2"

	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type BillboardsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewBillboardsRoutes(db *postgres.Postgres, l *logger.Logger) *BillboardsRoutes {
	return &BillboardsRoutes{
		db: db,
		l:  l,
	}
}

// Index godoc
// @Summary      Вывод постов для страницы афишы
// @Description  Настроенные по приоритету посты афишы
// @Tags         Клиент сторона
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200
// @Router       /v1/client/billboard [get].
func (p *BillboardsRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}
