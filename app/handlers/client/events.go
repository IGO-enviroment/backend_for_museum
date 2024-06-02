package handlers

import (
	"github.com/gofiber/fiber/v2"

	client_entity "museum/app/entity/client"
	"museum/app/handlers"
	"museum/app/models"
	client_repo "museum/app/repo/client"
	client_usecase "museum/app/usecase/client"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type ClientEventsRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewClientEventsRoutes(db *postgres.Postgres, l *logger.Logger) *ClientEventsRoutes {
	return &ClientEventsRoutes{
		db: db,
		l:  l,
	}
}

func (c *ClientEventsRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}

func (c *ClientEventsRoutes) Filter(ctx *fiber.Ctx) error {
	c.l.Info("NewEventsRepo")
	repo := client_repo.NewEventsRepo(c.db, c.l)
	entity := &client_entity.EventsEntity{}
	c.l.Info("EventsEntity")
	usecase := client_usecase.NewEventsCase(repo, entity)
	c.l.Info("NewEventsCase")
	result := usecase.Call()
	return ctx.Status(fiber.StatusAccepted).JSON(result)
}

type EventShow struct {
	ID int `json:"id"`
}

// Выдача конкретного мероприятия.
func (c *ClientEventsRoutes) Show(ctx *fiber.Ctx) error {
	var request EventShow
	if err := ctx.BodyParser(&request); err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неверные параметры запроса", nil),
		)
	}

	event, err := models.NewEventModel(c.db).Find(request.ID)
	if err != nil {
		c.l.Error(err)

		return ctx.Status(fiber.StatusBadRequest).JSON(
			handlers.NewErrorStruct("Неизвестная ошибка", nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(event)
}
