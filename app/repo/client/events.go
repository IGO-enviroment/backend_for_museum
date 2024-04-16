package repo

import (
	"context"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
)

type EventsRepo struct {
	db          *postgres.Postgres
	l           *logger.Logger
	eventsQuery squirrel.SelectBuilder
}

func NewEventsRepo(db *postgres.Postgres, l *logger.Logger) *EventsRepo {
	return &EventsRepo{
		db: db,
		l:  l,
	}
}

func (e *EventsRepo) AllEvents() {
	e.db.Builder.Select(e.selectFields()...).From("events")
}

func (e *EventsRepo) ByText(sql *squirrel.SelectBuilder) {
}

func (e *EventsRepo) RangeDate(start time.Time, end time.Time, sql *squirrel.SelectBuilder) {
	sql.Where("start")
}

func (e *EventsRepo) WithTickets(countTicket int, sql *squirrel.SelectBuilder) squirrel.SelectBuilder {
	return sql.Where(squirrel.Eq{"ticket_count": countTicket})
}

func (e *EventsRepo) WithArea(areaIds []int, sql *squirrel.SelectBuilder) {
	sql.InnerJoin("areas ON events.area_id=areas.id", squirrel.Eq{"areas.id": areaIds})
}

func (e *EventsRepo) WithType(typeIds []int, sql *squirrel.SelectBuilder) {
	sql.InnerJoin("type_events ON events.type_id=type_events.id", squirrel.Eq{"type_events.id": typeIds})
}

func (e *EventsRepo) WithTags(tagIds []int, sql *squirrel.SelectBuilder) {
	sql.LeftJoin(
		"event_tags ON event_tags.event_id = events.id",
		squirrel.Eq{"event_tags.id": tagIds},
	).LeftJoin(
		"tags ON event_tags.tag_id = tags.id",
	)
}

func (e *EventsRepo) WithPage(sql *squirrel.SelectBuilder, perPage, offset int) squirrel.SelectBuilder {
	return sql.Limit(uint64(perPage)).Offset(uint64(offset)).OrderBy("events.start_at ASC")
}

// Запрос на получения данных по событиям.
func (e *EventsRepo) GetValues(sql *squirrel.SelectBuilder) (pgx.Rows, bool) {
	query, args, err := sql.ToSql()
	if err != nil {
		e.l.Error(err)

		return nil, false
	}

	rows, err := e.db.Pool.Query(context.Background(), query, args...)
	if err != nil {
		e.l.Error(err)

		return nil, false
	}

	return rows, true
}

// Поля для выбора.
func (e *EventsRepo) selectFields() []string {
	return []string{
		"events.id", "events.title",
		"events.ticket_count", "events.start_at",
		"events.duration", "events.area_id",
		"events.type_id", "events.preview_url",
		"events.price", "events.created_at",
	}
}
