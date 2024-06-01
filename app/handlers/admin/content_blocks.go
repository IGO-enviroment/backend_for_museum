package handlers

import (
	contract_admin "museum/app/contract/admin"
	admin_entity "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
	usecase_admin "museum/app/usecase/admin"
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
	var request contract_admin.CreateContentBlocks
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.ValidatorErrors(err),
		)
	}

	usecase := usecase_admin.NewCreateContentBlocksCaseCase(
		admin_repo.NewContentBlocksRepo(c.db, c.l),
		&admin_entity.CreateContentBlocksEntity{
			ParentID:   request.ParentID,
			ParentType: request.ParentType,
			Type:       request.Type,
			Index:      request.Index,
			ValueStr:   request.ValueStr,
			ValueFile:  request.ValueFile,
		},
	)

	ok, err := usecase.Call()
	if !ok {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

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
