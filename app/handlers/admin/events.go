package handlers

import (
	contract_admin "museum/app/contract/admin"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
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
	var request contract_admin.CreateEvent
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err)

		return handlers.ErrorResponse(ctx)
	}

	repo := repo_admin.NewCreateEventRepo(e.db, e.l)
	usecase := usecase_admin.NewCreateEventCase(
		&repo,
		&entity_admin.CreateEventEntity{
			Title:       request.Title,
			Description: request.Description,
		},
	)
	result, err := usecase.Call()

	if err != nil {
		e.l.Error(err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": result,
	})
}

// Create godoc
func (e *EventsRoutes) Index(ctx *fiber.Ctx) error {
	usecase := usecase_admin.NewEventsTableCase(
		repo_admin.NewEventListRepo(e.db, e.l),
		&entity_admin.EventTable{},
	)
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

// Publish
func (e *EventsRoutes) Publish(ctx *fiber.Ctx) error {
	var request contract_admin.PublishEvent
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err)

		return handlers.ErrorResponse(ctx)
	}

	usecase := usecase_admin.NewPublishEventCase(
		repo_admin.NewPublishEventRepo(e.db, e.l),
		&entity_admin.PublishEventEntity{ID: request.ID},
	)

	ok, err := usecase.Call()
	if !ok {
		return ctx.Status(fiber.StatusOK).JSON(err)
	}

	return ctx.SendStatus(fiber.StatusCreated)
}
