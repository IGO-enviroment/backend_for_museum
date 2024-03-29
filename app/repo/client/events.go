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
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEventsRepo(db *postgres.Postgres, l *logger.Logger) *EventsRepo {
	return &EventsRepo{
		db: db,
		l:  l,
	}
}

func (e *EventsRepo) AllEvents() {
	e.db.Builder.Select("id").From("events")
}

func (e *EventsRepo) ByText(sql *squirrel.SelectBuilder) {
}

func (e *EventsRepo) RangeDate(start time.Time, end time.Time, sql *squirrel.SelectBuilder) {
	sql.Where("start")
}

func (e *EventsRepo) WithTickets(countTicket int, sql *squirrel.SelectBuilder) squirrel.SelectBuilder {
	return sql.Where(squirrel.Eq{"ticket_count": countTicket})
}

func (e *EventsRepo) WithAreas(areas []int, sql *squirrel.SelectBuilder) squirrel.SelectBuilder {
	withAreas := e.db.Builder.Select("id").From("areas").Where(squirrel.Eq{"id": areas})
	return sql.Where(squirrel.Eq{"area_id": withAreas})
}

func (e *EventsRepo) WithTypes(types []int, sql *squirrel.SelectBuilder) squirrel.SelectBuilder {
	withTypes := e.db.Builder.Select("id").From("type_events").Where(squirrel.Eq{"id": types})
	return sql.Where(squirrel.Eq{"type_id": withTypes})
}

func (e *EventsRepo) WithPage(sql *squirrel.SelectBuilder, perPage int, offset int) squirrel.SelectBuilder {
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
