package repo

import (
	"context"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type OptionsTicketFilter struct {
	TypesTicket []string
	Count       []string
	Price       []string
}

type EventsRepo struct {
	db    *postgres.Postgres
	Log   *logger.Logger
	Query squirrel.StatementBuilderType
}

func NewEventsRepo(db *postgres.Postgres, l *logger.Logger) *EventsRepo {
	return &EventsRepo{
		db:    db,
		Log:   l,
		Query: db.Builder,
	}
}

func (e *EventsRepo) ByText(sql *squirrel.SelectBuilder) {
}

func (e *EventsRepo) RangeDate(start time.Time, end time.Time, sql *squirrel.SelectBuilder) {
	sql.Where("start")
}

func (e *EventsRepo) WithArea(areaIds []int) {
	e.Query = e.Query.Where(squirrel.Eq{"events.area_id": areaIds})
}

func (e *EventsRepo) WithType(typeIds []int) {
	e.Query = e.Query.Where(squirrel.Eq{"events.type_id": typeIds})
}

func (e *EventsRepo) WithTags(tagIds []int) {
	tagsSql := e.db.Builder.Select("event_id").From("event_tags").Where(squirrel.Eq{"event_tags.tag_id": tagIds})
	e.Query = e.Query.Where(tagsSql.Prefix("events.id IN (").Suffix(")"))
}

func (e *EventsRepo) ByTicketData(options OptionsTicketFilter) {
	if len(options.TypesTicket) == 0 && len(options.Count) == 0 && len(options.Price) == 0 {
		return
	}

	var ticketSql squirrel.SelectBuilder
	var used bool

	eventIds := e.db.Builder.Select("event_id").From("tickets")

	if options.TypesTicket != nil && len(options.TypesTicket) > 0 {
		ticketSql = eventIds.Where(squirrel.Eq{"tickets.type": options.TypesTicket})
		used = true
	}

	if options.Count != nil {
		ticketSql = e.WithRangeValues(options.Count, "tickets.count", eventIds)
		used = true
	}

	if options.Price != nil {
		ticketSql = e.WithRangeValues(options.Count, "tickets.cost", eventIds)
		used = true
	}

	if !used {
		return
	}
	e.Query = e.Query.Where(ticketSql.Prefix("events.id IN (").Suffix(")"))
}

func (e *EventsRepo) WithRangeValues(values []string, column string, sql squirrel.SelectBuilder) squirrel.SelectBuilder {
	lenValues := len(values)

	if lenValues == 0 || lenValues > 2 {
		return sql
	}

	if len(values) == 1 {
		return sql.Where(squirrel.Eq{column: values[0]})
	}

	if values[1] != "" {
		sql = sql.Where(squirrel.LtOrEq{column: values})
	}

	if values[0] != "" {
		sql = sql.Where(squirrel.GtOrEq{column: values})
	}
	return sql
}

func (e *EventsRepo) CountValues() (int, bool) {
	var query string
	var args []interface{}
	var err error

	query, args, err = e.Query.Select("COUNT(id)").From("events").ToSql()

	if err != nil {
		e.Log.Error(err)

		return 0, false
	}

	return e.queryCount(query, args)
}

// Запрос на получения данных по событиям.
func (e *EventsRepo) GetValues(perPage, offset int) (pgx.Rows, bool) {
	query, args, err := e.withPage(
		e.Query.Select(e.selectFields()...).From("events"), perPage, offset,
	).ToSql()
	if err != nil {
		e.Log.Error(err)

		return nil, false
	}

	rows, err := e.db.Pool.Query(context.Background(), query, args...)
	if err != nil {
		e.Log.Error(err)

		return nil, false
	}

	return rows, true
}

// Поля для выбора.
func (e *EventsRepo) selectFields() []string {
	return []string{
		"events.id", "events.title", "events.start_at",
		"events.duration", "events.area_id",
		"events.type_id", "events.preview_url", "events.created_at",
	}
}

// Пагинация
func (e *EventsRepo) withPage(sql squirrel.SelectBuilder, perPage, offset int) squirrel.SelectBuilder {
	return sql.Limit(uint64(perPage)).Offset(uint64(offset)).OrderBy("events.start_at ASC")
}

// Запрос на количество записей
func (e *EventsRepo) queryCount(query string, args []interface{}) (int, bool) {
	var err error
	var count int

	err = e.db.Pool.QueryRow(context.Background(), query, args...).Scan(&count)
	if err != nil {
		e.Log.Error(err)

		return 0, false
	}

	return count, true
}
