package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

func ValidatorErrors(err error) *ErrorStruct {
	errorResponse := ErrorStruct{}

	validateErrors := err.(validator.ValidationErrors)
	if len(validateErrors) == 0 {
		errorResponse.Msg = "Неизвестная ошибка"

		return &errorResponse
	}

	for _, err := range validateErrors {
		errorResponse.Errors = append(errorResponse.Errors,
			ErrorWithKey{
				Key: err.Field(),
				Value: fmt.Sprintf(
					"Поле %s содержит недопустимое значение", err.Tag(),
				),
			},
		)
	}

	return &errorResponse
}
