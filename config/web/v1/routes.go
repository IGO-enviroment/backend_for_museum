package v1

import (
	handlers "museum/app/handlers"
	posts_handlers "museum/app/handlers/admin_handlers/posts"
	access_middleware "museum/app/handlers/middleware/access"
	auth_middleware "museum/app/handlers/middleware/auth"

	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	v1 := app.Group("/v1")

	// Авторизация
	auth := v1.Group("/auth")
	{
		auth_controller := handlers.NewAuthRoutes()
		auth.Post("/sign_in", auth_controller.SignIn)
		auth.Post("/sign_up", auth_controller.SignUp)
		auth.Get("/me", auth_middleware.AuthAccess, auth_controller.GetMe)
	}

	// Админская часть
	admin := v1.Group("/admin", auth_middleware.AuthAccess, access_middleware.AdminAccess)
	{
		// Работа с постами
		posts := admin.Group("/posts")

		posts_controller := posts_handlers.NewPostsRoutes()
		posts.Put("/update/:id/", posts_controller.Update)
		posts.Put("/:id/show", posts_controller.Show)
		posts.Get("/", posts_controller.Index)
		posts.Post("/", posts_controller.Create)

		// sales := admin.Group("/sales")

	}
}
