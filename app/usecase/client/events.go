package usecase

import (
	"fmt"
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

func (e *EventsCase) Filter() {
	sql := e.repo.Init()
	if e.entity.Tags != nil && len(e.entity.Tags) > 0 {
		sql = e.repo.WithTags(e.entity.Tags, &sql)
	}

	if e.entity.Types != nil && len(e.entity.Types) > 0 {
		sql = e.repo.WithType(e.entity.Types, &sql)
	}

	if e.entity.Areas != nil && len(e.entity.Areas) > 0 {
		sql = e.repo.WithArea(e.entity.Areas, &sql)
	}

	sql = e.repo.WithPage(&sql, 1, 1)
	rows, _ := e.repo.GetValues(&sql)
	fmt.Println(rows)
}
