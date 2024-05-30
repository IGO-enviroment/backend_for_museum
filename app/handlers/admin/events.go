package handlers

import (
	entity_admin "museum/app/entity/admin"
	repo_admin "museum/app/repo/admin"
	usecase_admin "museum/app/usecase/admin"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

type EventsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEventsRoutes(db *postgres.Postgres, l *logger.Logger) *EventsRoutes {
	return &EventsRoutes{
		db: db,
		l:  l,
	}
}

// Create godoc
func (e *EventsRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Create godoc
func (e *EventsRoutes) Index(ctx *fiber.Ctx) error {
	repo := repo_admin.NewEventListRepo(e.db, e.l)
	entity := &entity_admin.EventTable{}
	usecase := usecase_admin.NewEventsTableCase(repo, entity)
	result, err := usecase.Call()
	if err != nil {
		e.l.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}

// Create godoc
func (e *EventsRoutes) Show(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Publi
func (e *EventsRoutes) Publish(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}
