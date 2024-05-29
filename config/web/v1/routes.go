package v1

import (
	"github.com/gofiber/fiber/v2"
	admin_handlers "museum/app/handlers/admin"
	client_handlers "museum/app/handlers/client"
	"museum/app/middleware"
)

func SetRoutes(s *Server) {
	v1 := s.app.Group("/v1")

	ClientRoutes(s, v1)
	AdminsRoutes(s, v1)
}

// Сторона пользователя.
func ClientRoutes(s *Server, v1 fiber.Router) {
	client := v1.Group("/client")

	// Глобальный поиск по контенту
	{
		searchController := client_handlers.NewContentSearchRoutes(s.db, s.l)
		client.Get("/content/search", searchController.Search)
	}

	// Афиша (главная страница)
	{
		billboardsController := client_handlers.NewBillboardsRoutes(s.db, s.l)
		client.Get("/billboard", billboardsController.Index)
	}

	// Данные общего списка мероприятий
	{
		eventsController := client_handlers.NewClientEventsRoutes(s.db, s.l)
		client.Get("/index", eventsController.Index)
		client.Get("/events/filter", eventsController.Filter)
	}

	// Данные по популярным фильтрам
	{
		popularFilters := client_handlers.NewPopularFiltersRoutes(s.db, s.l)
		client.Get("/popular/filters", popularFilters.Index)
	}

	// Авторизация по почте
	{
		emailAuth := client_handlers.NewEmailAuthRoutes(s.db, s.l)
		client.Post("/sign_in", emailAuth.Create)
		client.Post("/auth", emailAuth.Verify)
	}
}

// Админская часть.
func AdminsRoutes(s *Server, v1 fiber.Router) {
	authPermissions := middleware.NewAuthAccess(s.db, s.l)

	admin := v1.Group(
		"/admin",
		authPermissions.Аuthorized, authPermissions.AdminAccess,
	)

	superAdmin := v1.Group(
		"/admin",
		authPermissions.Аuthorized, authPermissions.SuperAdminAccess,
	)

	auth := v1.Group("/auth")

	// Авторизация
	{
		authController := admin_handlers.NewAuthRoutes(s.db, s.l)

		// Супер админ
		register := superAdmin.Group("/register")
		register.Post("/", authController.AddUser)

		// Админ
		token := auth.Group("/token")
		token.Post("/generate", authController.GetToken)
	}

	// Работа с постами
	{
		posts := admin.Group("/posts")

		postsController := admin_handlers.NewPostsRoutes(s.db, s.l)
		posts.Put("/update/:id/", postsController.Update)
		posts.Put("/:id/show", postsController.Show)
		posts.Get("/", postsController.Index)
		posts.Post("/", postsController.Create)
	}

	// Темы событий, новостей и т.д.
	{
		topics := admin.Group("/topics")
		topicsController := admin_handlers.NewTopicRoutes(s.db, s.l)
		topics.Put("/update/:id", topicsController.Update)
		topics.Get("/", topicsController.Index)
		topics.Post("/", topicsController.Create)
	}

	// Типы мероприйтий
	{
		eventTypes := admin.Group("/event-types")
		eventTypesController := admin_handlers.NewEventTypesRoutes(s.db, s.l)
		eventTypes.Post("/create", eventTypesController.Create)
	}
}
