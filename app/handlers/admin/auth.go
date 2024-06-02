package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	admin_auth "museum/app/contract/admin/auth"
	"museum/app/contract/superadmin"
	"museum/app/handlers"
	"museum/app/models"
	admin_repo "museum/app/repo/admin"
	admin_usecase "museum/app/usecase/admin"
	"museum/app/utils"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"
)

type AuthRoutes struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewAuthRoutes(db *postgres.Postgres, l *logger.Logger) AuthRoutes {
	return AuthRoutes{db, l}
}

func (a *AuthRoutes) GetToken(ctx *fiber.Ctx) error {
	var loginModel admin_auth.LoginModel
	if err := ctx.BodyParser(&loginModel); err != nil {
		a.l.Error(err, "incorrect login model")

		return handlers.ErrorResponse(ctx)
	}

	getAdminCase := admin_usecase.NewGetAdminCase(admin_repo.NewAdminRepo(a.db, a.l), loginModel.Email)
	admin, err := getAdminCase.Call()
	if err != nil || admin == nil || !admin.IsAdmin || admin.Password != loginModel.Password {
		a.l.Error(err, "email or password incorrect")
		return handlers.ErrorResponse(ctx)
	}

	token, err := utils.GenerateToken(admin.Id, loginModel.Email, admin.Role)
	var twoMonth time.Duration = 1460
	cookie := &fiber.Cookie{
		Name:     "museum_client_auth",
		Value:    fmt.Sprintf("Bear %s", token),
		SameSite: "Lax",
		Secure:   false,
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * twoMonth).Local(),
	}
	ctx.Cookie(
		cookie,
	)

	return ctx.Status(fiber.StatusOK).JSON(cookie)
}

func (a *AuthRoutes) AddUser(ctx *fiber.Ctx) error {
	var loginModel superadmin.CreateAdminModel
	if err := ctx.BodyParser(&loginModel); err != nil {
		a.l.Error(err, "incorrect login model")

		return handlers.ErrorResponse(ctx)
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

func (a *AuthRoutes) Check(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("currentUser").(*models.User)
	if !ok {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}
