package admin

import (
	entity_admin "museum/app/entity/admin"
	admin_repo "museum/app/repo/admin"
)

type EventTypeCase struct {
	repo admin_repo.EventTypeRepo
}

func NewEventTypeUsecase(repo admin_repo.EventTypeRepo) EventTypeCase {
	return EventTypeCase{
		repo,
	}
}

func (c EventTypeCase) Create(entity entity_admin.EventTypeEntity) (int, error) {
	return c.repo.Create(entity)
}

func (c EventTypeCase) GetAll() ([]entity_admin.EventTypeEntity, error) {
	return c.repo.Get()
}
