package handlers

import (
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

type ContentBlocksRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewContentBlocksRoutes(db *postgres.Postgres, l *logger.Logger) ContentBlocksRoutes {
	return ContentBlocksRoutes{db, l}
}

// Создание блока контента.
func (c *ContentBlocksRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Обновление блока контента.
func (c *ContentBlocksRoutes) Update(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

// Выдача сохраненных блоков.
func (c *ContentBlocksRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}

// Удаление блока,
func (c *ContentBlocksRoutes) Delete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
