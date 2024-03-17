package handlers

import (
	"github.com/gofiber/fiber/v2"

	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type ClientEventsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewClientEventsRoutes(db *postgres.Postgres, l *logger.Logger) *ClientEventsRoutes {
	return &ClientEventsRoutes{
		db: db,
		l:  l,
	}
}

func (f *ClientEventsRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}

func (f *ClientEventsRoutes) Filter(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}
