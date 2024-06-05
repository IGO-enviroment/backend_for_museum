package handlers

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"
	"time"

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
	var areaEntity entity_admin.CreateAreaEntity
	if err := ctx.BodyParser(&areaEntity); err != nil {
		c.l.Error(err, "incorrect login model")
		return handlers.ErrorResponse(ctx)
	}

	if areaEntity.Longitude == 0 || areaEntity.Latitude == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Поля широта и долгота обязательны", nil),
		)
	}

	area := map[string]interface{}{
		"name":          areaEntity.Name,
		"description":   areaEntity.Description,
		"publish":       areaEntity.Publish,
		"address_value": areaEntity.AddressValue,
		"created_at":    time.Now(),
		"updated_at":    time.Now(),
		"longitude":     areaEntity.Longitude,
		"latitude":      areaEntity.Latitude,
	}

	sql, args, err := c.db.Builder.
		Insert("areas").
		SetMap(area).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
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
	defer rows.Close()

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			c.l.Error(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(
				handlers.NewErrorStruct("Неизвестная ошибка", nil),
			)
		}
		return ctx.Status(200).SendString(strconv.Itoa(id))
	}

	return ctx.Status(200).SendString(strconv.Itoa(0))
}

// Index areas.
func (c *AreasRoutes) Index(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.
		Select("id", "name", "description", "publish", "address_value", "created_at", "longitude", "latitude").
		From("areas").
		ToSql()
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
	defer rows.Close()

	var result []entity_admin.GetAreaEntity
	for rows.Next() {
		var area entity_admin.GetAreaEntity
		err = rows.Scan(
			&area.Id,
			&area.Name,
			&area.Description,
			&area.Publish,
			&area.AddressValue,
			&area.CreatedAt,
			&area.Longitude,
			&area.Latitude,
		)
		if err != nil {
			c.l.Error(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(
				handlers.NewErrorStruct("Неизвестная ошибка", nil),
			)
		}
		result = append(result, area)
	}

	return ctx.Status(fiber.StatusOK).JSON(result)
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

func (c *AreasRoutes) GetById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	sql, args, err := c.db.Builder.
		Select("id", "name", "description", "publish", "address_value", "created_at", "longitude", "latitude").
		From("areas").
		Where(squirrel.Eq{"id": id}).
		ToSql()
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
	defer rows.Close()

	if rows.Next() {
		var area entity_admin.GetAreaEntity
		err = rows.Scan(
			&area.Id,
			&area.Name,
			&area.Description,
			&area.Publish,
			&area.AddressValue,
			&area.CreatedAt,
			&area.Longitude,
			&area.Latitude,
		)
		if err != nil {
			c.l.Error(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(
				handlers.NewErrorStruct("Неизвестная ошибка", nil),
			)
		}
		return ctx.Status(fiber.StatusOK).JSON(area)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (c *AreasRoutes) DeleteById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	sql, args, err := c.db.Builder.
		Delete("areas").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	var pgError *pgconn.PgError
	_, err = c.db.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		errors.As(err, &pgError)
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct(pgError.Detail, nil),
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *AreasRoutes) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	var areaEntity entity_admin.CreateAreaEntity
	if err := ctx.BodyParser(&areaEntity); err != nil {
		c.l.Error(err, "incorrect login model")
		return handlers.ErrorResponse(ctx)
	}

	if areaEntity.Longitude == 0 || areaEntity.Latitude == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Поля широта и долгота обязательны", nil),
		)
	}
	sql, args, err := c.db.Builder.
		Update("areas").
		Set("name", areaEntity.Name).
		Set("description", areaEntity.Description).
		Set("publish", areaEntity.Publish).
		Set("address_value", areaEntity.AddressValue).
		Set("longitude", areaEntity.Longitude).
		Set("latitude", areaEntity.Latitude).
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Unable to build UPDATE query", nil),
		)
	}

	var pgError *pgconn.PgError
	_, err = c.db.Pool.Exec(context.Background(), sql, args...)
	if err != nil {
		errors.As(err, &pgError)
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct(pgError.Detail, nil),
		)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
