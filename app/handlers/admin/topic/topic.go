// Статистика по продажам билетов
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type TopicRoutes struct {
	layout string
}

func NewTopicRoutes() *TopicRoutes {
	return &TopicRoutes{
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
// @Router       /v1/admin/topics [post]
func (p *TopicRoutes) Create(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Update godoc
// @Summary      Обновление данных поста
// @Description  Создание нового поста
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /v1/admin/topics/update/{id} [put]
func (p *TopicRoutes) Update(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusCreated)
}

// Index godoc
// @Summary      Все темы
// @Description  Вывод списка всех доступных тем для событий и т.д.
// @Tags         Админка
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Router       /v1/admin/topics [get]
func (p *TopicRoutes) Index(ctx *fiber.Ctx) error {
	allTopics := []string{"123", "123", ""}
	return ctx.Render(
		"admin/topics/index",
		fiber.Map{
			"Topics": allTopics,
		},
		p.layout,
	)
}
