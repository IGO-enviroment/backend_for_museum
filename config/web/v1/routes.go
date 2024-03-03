package v1

import (
	handlers "museum/app/handlers"
	posts_handlers "museum/app/handlers/admin/posts"
	topics_handelrs "museum/app/handlers/admin/topic"
	"museum/app/middleware"
)

func SetRoutes(s *Server) {
	v1 := s.app.Group("/v1")

	authPermissions := middleware.NewAuthAccess(s.db, s.l)

	// Авторизация
	auth := v1.Group("/auth")
	{
		authController := handlers.NewAuthRoutes(s.db, s.l)
		auth.Post("/sign_in", authController.SignIn)
		auth.Post("/sign_up", authController.SignUp)
		auth.Get("/me", authPermissions.Аuthorized, authController.GetMe)
	}

	// Проверка билетов
	// v1.Get("/verify/info/:code")

	// Админская часть
	admin := v1.Group("/admin", authPermissions.Аuthorized, authPermissions.AdminAccess)
	{
		// sales := admin.Group("/sales")
		// Работа с постами
		posts := admin.Group("/posts")
		{
			postsController := posts_handlers.NewPostsRoutes()
			posts.Put("/update/:id/", postsController.Update)
			posts.Put("/:id/show", postsController.Show)
			posts.Get("/", postsController.Index)
			posts.Post("/", postsController.Create)
		}

		// Темы событий, новостей и т.д.
		topics := admin.Group("/topics")
		{
			topicsController := topics_handelrs.NewTopicRoutes()
			topics.Put("/update/:id", topicsController.Update)
			topics.Get("/", topicsController.Index)
			topics.Post("/", topicsController.Create)
		}
	}
}
