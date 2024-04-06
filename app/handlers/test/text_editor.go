package handlers

import (
	"github.com/gofiber/fiber/v2"
	entity_test "museum/app/entity/test"
	"museum/app/handlers"
	test_repo "museum/app/repo/test"
	test_usecase "museum/app/usecase/test"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"strconv"
)

type TestTextEditorRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewTestTextEditorRoutes(db *postgres.Postgres, l *logger.Logger) *TestTextEditorRoutes {
	return &TestTextEditorRoutes{
		db: db,
		l:  l,
	}
}

func (t TestTextEditorRoutes) Create(ctx *fiber.Ctx) error {
	var request *string
	if err := ctx.BodyParser(&request); err != nil {
		t.l.Error(err, "http - v1 - doTranslate")
		return handlers.ErrorResponse(ctx)
	}

	testTextEditor := entity_test.TestTextEditorContent{
		TypeValue:  "text",
		DataValue:  *request,
		OrderValue: 0,
		ModelID:    0,
		ModelType:  "test",
		Options:    "{}",
	}

	usercase := test_usecase.NewTestTextEditorCreateCase(test_repo.NewTextEditorRepo(t.db, t.l), &testTextEditor)
	id, err := usercase.Call()
	if err != nil {
		return handlers.ErrorResponse(ctx)
	}
	return ctx.SendString(strconv.Itoa(id))
}

func (t TestTextEditorRoutes) Get(ctx *fiber.Ctx) error {
	contentIdStr := ctx.Params("id")
	contentId, err := strconv.Atoi(contentIdStr)
	if err != nil {
		return handlers.ErrorResponse(ctx)
	}
	usercase := test_usecase.NewTestTextEditorGetCase(test_repo.NewTextEditorRepo(t.db, t.l), contentId)
	content, err := usercase.Call()
	if err != nil {
		return handlers.ErrorResponse(ctx)
	}
	if content == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}
	return ctx.Status(fiber.StatusAccepted).JSON(content)
}

func (t TestTextEditorRoutes) Update(ctx *fiber.Ctx) error {
	var request *string
	if err := ctx.BodyParser(&request); err != nil {
		t.l.Error(err, "http - v1 - doTranslate")
		return handlers.ErrorResponse(ctx)
	}
	contentIdStr := ctx.Params("id")
	contentId, err := strconv.Atoi(contentIdStr)
	if err != nil {
		return handlers.ErrorResponse(ctx)
	}
	usercase := test_usecase.NewTestTextEditorUpdateCase(test_repo.NewTextEditorRepo(t.db, t.l), contentId, *request)
	err = usercase.Call()
	if err != nil {
		return handlers.ErrorResponse(ctx)
	}
	return ctx.SendStatus(fiber.StatusAccepted)
}
