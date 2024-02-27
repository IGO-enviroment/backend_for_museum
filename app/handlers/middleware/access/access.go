/*
 * Промежуточные проверки перед обратокой контроллеров
 */
package middleware

import (
	users "museum/app/models/user"

	"github.com/gofiber/fiber/v2"
)

// Проверка на наличие роли админа
func AdminAccess(ctx *fiber.Ctx) error {
	user := ctx.Locals("currentUser").(*users.User)
	ok := user.Query.IsAdmin()
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	return ctx.Next()
}
