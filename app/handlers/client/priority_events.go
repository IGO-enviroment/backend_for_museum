package handlers

import (
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

type PriorityEventsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewPriorityEvents(db *postgres.Postgres, l *logger.Logger) *PriorityEventsRoutes {
	return &PriorityEventsRoutes{
		db: db,
		l:  l,
	}
}

func (p *PriorityEventsRoutes) Get(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}
