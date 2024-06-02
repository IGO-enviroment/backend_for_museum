package handlers

import (
	"github.com/gofiber/fiber/v2"
	"museum/app/contract/admin"
	admin_entity "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
	admin_usecase "museum/app/usecase/admin"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"
)

type EventTypesRoutes struct {
	db            *postgres.Postgres
	l             *logger.Logger
	eventTypeCase admin_usecase.EventTypeCase
}

func NewEventTypesRoutes(db *postgres.Postgres, l *logger.Logger) EventTypesRoutes {
	return EventTypesRoutes{
		db:            db,
		l:             l,
		eventTypeCase: admin_usecase.NewEventTypeUsecase(admin_repo.NewEventTypeRepo(db, l)),
	}
}

func (e *EventTypesRoutes) Create(ctx *fiber.Ctx) error {
	var request admin.CreateEventType
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	eventType := admin_entity.EventTypeEntity{
		Name:        request.Name,
		Description: request.Description,
		IsVisible:   request.IsVisible,
	}
	eventId, err := e.eventTypeCase.Create(eventType)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendString(strconv.Itoa(eventId))
}

func (e *EventTypesRoutes) GetAll(ctx *fiber.Ctx) error {
	result, err := e.eventTypeCase.GetAll()
	if err != nil {
		e.l.Error(err)

		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.Status(fiber.StatusOK).JSON(result)
}
