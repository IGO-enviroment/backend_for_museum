package handlers

import (
	"context"
	"museum/app/contract/admin"
	admin_entity "museum/app/entity/admin"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
	admin_usecase "museum/app/usecase/admin"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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

// Выдача типов мероприятий для создание меры.
func (c *EventTypesRoutes) IndexEventTypesID(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.Select("id", "name").From("type_events").ToSql()
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	rows, err := c.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	tags, err := pgx.CollectRows(
		rows, pgx.RowToStructByName[entity_admin.EventTypeIDSEntity],
	)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}
