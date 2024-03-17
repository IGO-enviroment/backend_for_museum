// Статистика по продажам билетов
package handlers

import "github.com/gofiber/fiber/v2"

type PostsRoutes struct {
	layout string
}

func NewPostsRoutes() *PostsRoutes {
	return &PostsRoutes{
		layout: "layouts/admin",
	}
}

// Create godoc
// @Summary      Создание новой новости
// @Description  Создание нового поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/posts [post].
func (p *PostsRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Update godoc
// @Summary      Обновление данных поста
// @Description  Создание нового поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/admin/posts/update/{id} [put].
func (p *PostsRoutes) Update(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Index godoc
// @Summary      Все посты
// @Description  Вывод таблицы по всем постам
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/posts [get].
func (p *PostsRoutes) Index(ctx *fiber.Ctx) error {
	allPosts := []string{"123", "123", ""}

	return ctx.Render(
		"admin/posts/index",
		fiber.Map{
			"Posts": allPosts,
		},
		p.layout,
	)
}

// Show godoc
// @Summary      Отображение конкретного поста
// @Description  Вывод одного поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/posts/{id}/show [get].
func (p *PostsRoutes) Show(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}
