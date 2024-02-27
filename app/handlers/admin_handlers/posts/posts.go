/*
 * Статистика по продажам билетов
 */
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type postsRoutes struct{}

func NewPostsRoutes() *postsRoutes {
	return &postsRoutes{}
}

// Create godoc
// @Summary      Создание новой новости
// @Description  Создание нового поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/posts [post]
func (p *postsRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Update godoc
// @Summary      Обновление данных поста
// @Description  Создание нового поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/admin/posts/update/{id} [put]
func (p *postsRoutes) Update(ctx *fiber.Ctx) error {
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
// @Router       /v1/admin/posts [get]
func (p *postsRoutes) Index(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Show godoc
// @Summary      Отображение конкретного поста
// @Description  Вывод одного поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/posts/{id}/show [get]
func (p *postsRoutes) Show(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}
