/*
 * Статистика по продажам билетов
 */
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type postPublishRoutes struct{}

func NewPostPublishRoutes() *postPublishRoutes {
	return &postPublishRoutes{}
}

// Create godoc
// @Summary      Публикация поста
// @Description  Отображение поста для всех пользователей
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/admin/posts/publish/{id} [post]
func (p *postPublishRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Delete godoc
// @Summary      Убрать статью с публикации
// @Description  Статья больше не будет отображаться для пользователей
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/admin/posts/publish/{id} [post]
func (p *postPublishRoutes) Delete(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}
