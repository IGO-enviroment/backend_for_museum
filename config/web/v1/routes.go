package v1

import (
	admin_handlers "museum/app/handlers/admin"
	client_handlers "museum/app/handlers/client"
	"museum/app/middleware"

	"github.com/gofiber/fiber/v2"
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

	checkAuth := v1.Group("/check", authPermissions.Аuthorized)

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

		// Проверка на валидность токена
		checkAuth.Get("/", authController.Check)
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
		eventTypes.Get("/", eventTypesController.GetAll)
		eventTypes.Delete("/:id", eventTypesController.Delete)
		eventTypes.Put("/:id", eventTypesController.Update)
	}

	// Площадки
	{
		areas := admin.Group("/areas")
		areasController := admin_handlers.NewAreasRoutes(s.db, s.l)
		areas.Post("/create", areasController.Create)
		areas.Get("/", areasController.Index)
		areas.Get("/:id", areasController.GetById)
		areas.Delete("/:id", areasController.DeleteById)
		areas.Put("/:id", areasController.Update)
	}

	EventsRoutes(s, admin)

	ContentBlocksRoutes(s, admin)
}

// Мероприятия..
func EventsRoutes(s *Server, admin fiber.Router) {
	events := admin.Group("/events")
	eventsController := admin_handlers.NewEventsRoutes(s.db, s.l)
	events.Post("/", eventsController.Create)
	events.Get("/show/:id", eventsController.Show)
	events.Get("/", eventsController.Index)
	events.Put("/publish/:id", eventsController.Publish)
}

// Контент блоки.
func ContentBlocksRoutes(s *Server, admin fiber.Router) {
	contentBlocks := admin.Group("/content-block")
	contentBlockController := admin_handlers.NewContentBlocksRoutes(s.db, s.l)

	contentBlocks.Get("/index", contentBlockController.Index)
	contentBlocks.Post("/create", contentBlockController.Create)
	contentBlocks.Put("/update/:id", contentBlockController.Update)
	contentBlocks.Delete("/delete/:id", contentBlockController.Delete)
}
