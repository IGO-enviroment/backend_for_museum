package handlers

import (
	"github.com/gofiber/fiber/v2"

	client_entity "museum/app/entity/client"
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

func (f *ClientEventsRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusAccepted)
}

func (f *ClientEventsRoutes) Filter(ctx *fiber.Ctx) error {
	f.l.Info("NewEventsRepo")
	repo := client_repo.NewEventsRepo(f.db, f.l)
	entity := &client_entity.EventsEntity{}
	f.l.Info("EventsEntity")
	usecase := client_usecase.NewEventsCase(repo, entity)
	f.l.Info("NewEventsCase")
	result := usecase.Call()
	return ctx.Status(fiber.StatusAccepted).JSON(result)
}
