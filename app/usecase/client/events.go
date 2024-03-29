package usecase

import (
	entity_client "museum/app/entity/client"
	repo_client "museum/app/repo/client"
)

type EventsCase struct {
	repo   *repo_client.EventsRepo
	entity *entity_client.EventsEntity
}

func NewEventsCase(repo *repo_client.EventsRepo, entity *entity_client.EventsEntity) *EventsCase {
	return &EventsCase{
		repo:   repo,
		entity: entity,
	}
}

func (e *EventsCase) Filter() {}
