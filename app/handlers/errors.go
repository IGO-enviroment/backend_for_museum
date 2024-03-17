package handlers

import "github.com/gofiber/fiber/v2"

type ErrorWithKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorStruct struct {
	Msg    string         `json:"message"`
	Errors []ErrorWithKey `json:"errors"`
}

func ErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ErrorStruct{})
}
