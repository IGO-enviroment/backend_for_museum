package handlers

import (
	"github.com/gofiber/fiber/v2"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type EventTypesRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEventTypesRoutes(db *postgres.Postgres, l *logger.Logger) EventTypesRoutes {
	return EventTypesRoutes{db, l}
}

func (p *EventTypesRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}
