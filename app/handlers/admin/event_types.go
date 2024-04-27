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
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEventTypesRoutes(db *postgres.Postgres, l *logger.Logger) EventTypesRoutes {
	return EventTypesRoutes{db, l}
}

func (e *EventTypesRoutes) Create(ctx *fiber.Ctx) error {
	var request admin.CreateEventType
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	adminCase := admin_usecase.NewCreateEventType(
		admin_repo.NewEventTypeRepo(e.db, e.l),
		admin_entity.EventTypeEntity{
			Name:        request.Name,
			Description: request.Description,
			IsVisible:   request.IsVisible,
		})
	eventId, err := adminCase.Call()
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendString(strconv.Itoa(eventId))
}
