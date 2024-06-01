package usecase

import (
	"fmt"
	entity_client "museum/app/entity/client"
	repo_client "museum/app/repo/client"

	"github.com/jackc/pgx/v5"
)

const perPage = 10

type EventsCase struct {
	repo     *repo_client.EventsRepo
	entity   *entity_client.EventsEntity
	response entity_client.EventsResponse
}

func NewEventsCase(repo *repo_client.EventsRepo, entity *entity_client.EventsEntity) *EventsCase {
	return &EventsCase{
		repo:     repo,
		entity:   entity,
		response: entity_client.EventsResponse{},
	}
}

func (e *EventsCase) Call() entity_client.EventsResponse {
	var countRows int
	var ok bool

	if e.entity.Tags != nil && len(e.entity.Tags) > 0 {
		e.repo.WithTags(e.entity.Tags)
	}

	if e.entity.Types != nil && len(e.entity.Types) > 0 {
		e.repo.WithType(e.entity.Types)
	}

	if e.entity.Areas != nil && len(e.entity.Areas) > 0 {
		e.repo.WithArea(e.entity.Areas)
	}

	e.repo.ByTicketData(
		repo_client.OptionsTicketFilter{
			Price:       e.entity.Price,
			Count:       e.entity.TicketCount,
			TypesTicket: e.entity.TypeTicket,
		},
	)

	e.repo.Log.Info("Count")

	countRows, ok = e.repo.CountValues()
	if !ok {
		return e.response
	}

	e.setPageData(countRows)

	fmt.Println(countRows)
	fmt.Println(ok)
	rows, ok := e.eventRows()
	if !ok {
		return e.response
	}

	e.setData(rows)

	e.repo.Log.Info("response")

	return e.response
}

func (e *EventsCase) eventRows() (pgx.Rows, bool) {
	if e.entity.Page == nil {
		return e.repo.GetValues(perPage, 1)
	}

	return e.repo.GetValues(perPage, *e.entity.Page)
}

func (e *EventsCase) setPageData(countRows int) {
	e.response.Page.Total = countRows
	e.response.Page.Current = 1
}

func (e *EventsCase) setData(rows pgx.Rows) {
	for rows.Next() {
		var fff entity_client.EventItem

		err := rows.Scan(&fff)
		if err != nil {
			e.repo.Log.Info(err.Error())
		}

		e.repo.Log.Info(fmt.Sprintf("%#v", fff))
	}
}
