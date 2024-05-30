package admin

import (
	"context"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/jackc/pgx/v4"
)

type EventListRepo struct {
	db     *postgres.Postgres
	l      *logger.Logger
	result pgx.Rows
}

func NewEventListRepo(db *postgres.Postgres, l *logger.Logger) *EventListRepo {
	return &EventListRepo{
		db: db,
		l:  l,
	}
}

func (e *EventListRepo) Call() (pgx.Rows, error) {
	err := e.query()
	if err != nil {
		return nil, err
	}

	return e.result, nil
}

func (e *EventListRepo) query() error {
	sql, args, err := e.db.Builder.Select(
		"events.id, events.title, events.publish, events.ticket_count",
		"type_events.name as type, areas.name as area, events.created_at",
	).From("events").LeftJoin(
		"type_events ON type_events.id = events.type_id",
	).LeftJoin(
		"areas ON areas.id = events.area_id",
	).ToSql()
	if err != nil {
		return err
	}

	rows, err := e.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return err
	}

	e.result = rows

	return nil
}
