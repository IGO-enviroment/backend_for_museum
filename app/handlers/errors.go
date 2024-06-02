package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

type ErrorWithKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ErrorStruct struct {
	Msg    string         `json:"message"`
	Errors []ErrorWithKey `json:"errors"`
}

func NewErrorStruct(msg string, fields *map[string]string) *ErrorStruct {
	response := ErrorStruct{Msg: msg, Errors: []ErrorWithKey{}}

	if fields != nil {
		for key, value := range *fields {
			response.Errors = append(response.Errors, ErrorWithKey{Key: key, Value: value})
		}
	}

	return &response
}

func ValidatorErrors(err validate.Errors) *ErrorStruct {
	response := ErrorStruct{Msg: "", Errors: []ErrorWithKey{}}

	for field, _ := range err.All() {
		response.Errors = append(response.Errors, ErrorWithKey{
			Key: field, Value: err.FieldOne(field),
		})
	}

	return &response
}

func ErrorResponse(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(ErrorStruct{})
}
