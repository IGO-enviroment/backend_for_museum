package handlers

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
	"museum/app/contract/admin"
	admin_entity "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
	admin_usecase "museum/app/usecase/admin"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"
	"time"
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

func (e *EventTypesRoutes) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		e.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}

	var request admin.CreateEventType
	if err := ctx.BodyParser(&request); err != nil {
		e.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	sql, args, err := e.db.Builder.
		Update("type_events").
		Set("name", request.Name).
		Set("description", request.Description).
		Set("is_visible", request.IsVisible).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		e.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Unable to build UPDATE query", nil),
		)
	}

	var pgError *pgconn.PgError
	_, err = e.db.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		errors.As(err, &pgError)
		e.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct(pgError.Detail, nil),
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (e *EventTypesRoutes) Delete(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		e.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	sql, args, err := e.db.Builder.
		Delete("type_events").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	var pgError *pgconn.PgError
	_, err = e.db.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		errors.As(err, &pgError)
		e.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct(pgError.Detail, nil),
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
