package admin

import (
	entity_admin "museum/app/entity/admin"
	admin_repo "museum/app/repo/admin"
)

type CreateEventTypeCase struct {
	repo   admin_repo.EventTypeRepo
	entity entity_admin.EventTypeEntity
}

func NewCreateEventType(
	repo admin_repo.EventTypeRepo,
	entity entity_admin.EventTypeEntity) CreateEventTypeCase {
	return CreateEventTypeCase{
		repo,
		entity,
	}
}

func (c CreateEventTypeCase) Call() (int, error) {
	return c.repo.Create(c.entity)
}
