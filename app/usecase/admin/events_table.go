package admin

import (
	entity_admin "museum/app/entity/admin"
	repo_admin "museum/app/repo/admin"

	"github.com/jackc/pgx/v5"
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

// Выгружаем все события.
func (e *EventsTableCase) Call() (entity_admin.EventTable, error) {
	rows, err := e.repo.Call()
	if err != nil {
		return entity_admin.EventTable{}, err
	}

	events, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity_admin.EventForTable])
	if err != nil {
		return entity_admin.EventTable{}, err
	}

	result := entity_admin.EventTable{
		Events: events,
	}

	return result, nil
}
