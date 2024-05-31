package admin

import (
	entity_admin "museum/app/entity/admin"
	repo_admin "museum/app/repo/admin"
)

type EventsTableCase struct {
	repo   *repo_admin.EventListRepo
	entity *entity_admin.EventTable
}

func NewEventsTableCase(
	repo *repo_admin.EventListRepo,
	entity *entity_admin.EventTable) EventsTableCase {
	return EventsTableCase{
		repo:   repo,
		entity: entity,
	}
}

func (e *EventsTableCase) Call() (entity_admin.EventTable, error) {
	rows, err := e.repo.Call()
	if err != nil {
		return entity_admin.EventTable{}, err
	}

	return e.entity.ScanFromEquery(rows)
}
