package usecase

import (
	"fmt"
	entity_client "museum/app/entity/client"
	repo_client "museum/app/repo/client"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

const perPage = 10

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

func (e *EventsCase) FilterEvents() {
	var sql squirrel.StatementBuilderType

	if e.entity.Tags != nil && len(e.entity.Tags) > 0 {
		sql = e.repo.WithTags(e.entity.Tags, &sql)
	}

	if e.entity.Types != nil && len(e.entity.Types) > 0 {
		sql = e.repo.WithType(e.entity.Types, &sql)
	}

	if e.entity.Areas != nil && len(e.entity.Areas) > 0 {
		sql = e.repo.WithArea(e.entity.Areas, &sql)
	}

	sql = e.repo.ByTicketData(
		repo_client.OptionsTicketFilter{
			Price:       e.entity.Price,
			Count:       e.entity.TicketCount,
			TypesTicket: e.entity.TypeTicket,
		},
		sql,
	)

	countRows, ok := e.repo.Count(&sql)
	fmt.Println(countRows)
	fmt.Println(ok)
	e.EventRows(sql)
}

func (e *EventsCase) EventRows(sql squirrel.StatementBuilderType) (pgx.Rows, bool) {
	var query squirrel.SelectBuilder

	selectQuery := e.repo.AllEvents(&sql)

	if e.entity.Page == nil {
		query = e.repo.WithPage(&selectQuery, perPage, 1)
	} else {
		query = e.repo.WithPage(&selectQuery, perPage, *e.entity.Page)
	}

	return e.repo.GetValues(&query)
}
