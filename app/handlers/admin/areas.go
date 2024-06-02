package handlers

import (
	"context"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type AreasRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewAreasRoutes(db *postgres.Postgres, l *logger.Logger) *AreasRoutes {
	return &AreasRoutes{
		db: db,
		l:  l,
	}
}

// Create areas.
func (c *AreasRoutes) Create(ctx *fiber.Ctx) error {
	return nil
}

// Index areas.
func (c *AreasRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (c *AreasRoutes) IndexAreasID(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.Select("id", "name").From("areas").ToSql()
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
		rows, pgx.RowToStructByName[entity_admin.AreaIDSEntity],
	)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}
