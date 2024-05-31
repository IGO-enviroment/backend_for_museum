package handlers

import (
	"github.com/gofiber/fiber/v2"

	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type PopularFiltersRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewPopularFiltersRoutes(db *postgres.Postgres, l *logger.Logger) *PopularFiltersRoutes {
	return &PopularFiltersRoutes{
		db: db,
		l:  l,
	}
}

// Index godoc
// @Summary      Список популярных фильтров
// @Description  Настроенные админами популярные фильтры
// @Tags         Клиент сторона
// @Accept       json
// @Produce      json
// @Success    	 200
// @Router       /v1/client/popular/filters [get].
func (p *PopularFiltersRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}
