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

type TagsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewTagsRoutes(db *postgres.Postgres, l *logger.Logger) *TagsRoutes {
	return &TagsRoutes{
		db: db,
		l:  l,
	}
}

// Create areas.
func (c *TagsRoutes) Create(ctx *fiber.Ctx) error {
	var areaEntity entity_admin.CreateTagEntity
	if err := ctx.BodyParser(&areaEntity); err != nil {
		c.l.Error(err, "incorrect login model")
		return handlers.ErrorResponse(ctx)
	}

	area := map[string]interface{}{
		"name":        areaEntity.Name,
		"description": areaEntity.Description,
		"group_name":  areaEntity.GroupName,
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
	}

	sql, args, err := c.db.Builder.
		Insert("tags").
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
func (c *TagsRoutes) Index(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.Select("id", "name", "description", "group_name", "created_at").From("tags").ToSql()
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
		rows, pgx.RowToStructByName[entity_admin.TagEntity],
	)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}

// Выдача тегов для мероприятия.
func (c *TagsRoutes) IndexTagsID(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.Select("id", "name").From("tags").ToSql()
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
		rows, pgx.RowToStructByName[entity_admin.TagsIDSEntity],
	)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}

func (c *TagsRoutes) GetGroups(ctx *fiber.Ctx) error {
	sql, args, err := c.db.Builder.Select("group_name").Distinct().From("tags").ToSql()
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
		rows, pgx.RowTo[string],
	)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(tags)
}

func (c *TagsRoutes) GetById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	sql, args, err := c.db.Builder.
		Select("id", "name", "description", "group_name", "created_at", "updated_at").
		From("tags").
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

	tags, err := pgx.CollectRows(
		rows, pgx.RowToStructByName[entity_admin.TagEntity],
	)
	if len(tags) > 0 {
		return ctx.Status(fiber.StatusOK).JSON(tags[0])
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{})
}

func (c *TagsRoutes) DeleteById(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	sql, args, err := c.db.Builder.
		Delete("tags").
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

func (c *TagsRoutes) Update(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.l.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Некорректный id", nil),
		)
	}
	var tagEntity entity_admin.CreateTagEntity
	if err := ctx.BodyParser(&tagEntity); err != nil {
		c.l.Error(err, "incorrect tag model")
		return handlers.ErrorResponse(ctx)
	}

	sql, args, err := c.db.Builder.
		Update("tags").
		Set("name", tagEntity.Name).
		Set("description", tagEntity.Description).
		Set("group_name", tagEntity.GroupName).
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
