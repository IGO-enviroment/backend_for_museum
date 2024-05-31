package admin

import (
	entity_admin "museum/app/entity/admin"
	"museum/app/models"
	repo_admin "museum/app/repo/admin"
)

type CreateEventCase struct {
	repo   *repo_admin.CreateEventRepo
	entity *entity_admin.CreateEventEntity
}

func NewCreateEventCase(repo *repo_admin.CreateEventRepo, entity *entity_admin.CreateEventEntity) *CreateEventCase {
	return &CreateEventCase{
		repo:   repo,
		entity: entity,
	}
}

func (e *CreateEventCase) Call() (int, error) {
	id, err := e.repo.Call(e.CollectData())
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (e *CreateEventCase) CollectData() *models.Event {
	data := models.Event{
		Title:   e.entity.Title,
		StartAt: e.entity.StartAt,
	}

	return &data
}
