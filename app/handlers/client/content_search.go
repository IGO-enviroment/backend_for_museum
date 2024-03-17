package handlers

import (
	client_contract "museum/app/contract/client"
	client_entity "museum/app/entity/client"
	"museum/app/handlers"
	client_repo "museum/app/repo/client"
	client_usecase "museum/app/usecase/client"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/gofiber/fiber/v2"
)

type ContentSearchRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewContentSearchRoutes(db *postgres.Postgres, l *logger.Logger) *ContentSearchRoutes {
	return &ContentSearchRoutes{
		db: db,
		l:  l,
	}
}

// Search godoc
// @Summary      Поиск по названиям и контенту
// @Description  Поиск мероприятий, новостей, статей по тексту и контенту внутри них
// @Tags         Клиент сторона
// @Accept       json
// @Produce      json
// @Success    	 200 {object} client_entity.ResultContentSearch
// @Failure    	 422 {object} handlers.ErrorStruct
// @Router       /v1/client/content/search [get].
func (c *ContentSearchRoutes) Search(ctx *fiber.Ctx) error {
	var request client_contract.ContentSearch
	if err := ctx.BodyParser(&request); err != nil {
		c.l.Error(err, "http - v1 - doTranslate")

		return handlers.ErrorResponse(ctx)
	}

	usercase := client_usecase.NewContentSearch(
		client_repo.NewContentSearchRepo(c.db, c.l),
		&client_entity.SearchEntity{TypeSearch: "all", Target: "target"},
	)
	result := usercase.Call()

	return ctx.Status(fiber.StatusAccepted).JSON(result)
}
