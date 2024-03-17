package v1

import (
	admin_handlers "museum/app/handlers/admin"
	client_handlers "museum/app/handlers/client"
	"museum/app/middleware"
)

func SetRoutes(s *Server) {
	v1 := s.app.Group("/v1")

	authPermissions := middleware.NewAuthAccess(s.db, s.l)

	// Сторона пользователя
	client := v1.Group("/client")
	{
		// Глобальный поиск по контенту
		searchController := client_handlers.NewContentSearchRoutes(s.db, s.l)
		client.Get("/content/search", searchController.Search)

		// Афиша (главная страница)
		billboardsController := client_handlers.NewBillboardsRoutes(s.db, s.l)
		client.Get("/billboard", billboardsController.Index)

		// Данные общего списка мероприятий
		eventsController := client_handlers.NewClientEventsRoutes(s.db, s.l)
		client.Get("/index", eventsController.Index)
		client.Get("/filter", eventsController.Filter)

		// Данные по популярным фильтрам
		popularFilters := client_handlers.NewPopularFiltersRoutes(s.db, s.l)
		client.Get("/popular/filters", popularFilters.Index)
	}

	// Админская часть
	admin := v1.Group(
		"/admin",
		authPermissions.Аuthorized, authPermissions.AdminAccess,
	)
	{
		// Работа с постами
		posts := admin.Group("/posts")
		{
			postsController := admin_handlers.NewPostsRoutes()
			posts.Put("/update/:id/", postsController.Update)
			posts.Put("/:id/show", postsController.Show)
			posts.Get("/", postsController.Index)
			posts.Post("/", postsController.Create)
		}

		// Темы событий, новостей и т.д.
		topics := admin.Group("/topics")
		{
			topicsController := admin_handlers.NewTopicRoutes()
			topics.Put("/update/:id", topicsController.Update)
			topics.Get("/", topicsController.Index)
			topics.Post("/", topicsController.Create)
		}
	}
}
