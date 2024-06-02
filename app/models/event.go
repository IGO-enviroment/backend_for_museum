package models

import (
	"context"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// Мероприятия.
type Event struct {
	ID          int
	Title       string
	Description string

	Publish bool

	TicketCount int

	StartAt time.Time

	Duration int

	AreaID int
	TypeID int
	TagIDS []int

	PreviewURL string

	CreatedAt time.Time
	UpdateAt  time.Time
}

type EventModel struct {
	event Event
	db    *postgres.Postgres
}

func NewEventModel(db *postgres.Postgres) *EventModel {
	return &EventModel{
		db: db,
	}
}

func (e *EventModel) Find(id int) (Event, error) {
	sql, args, err := e.db.Builder.Select("*").From(
		"events",
	).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return e.event, err
	}

	rows, err := e.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return e.event, err
	}

	e.event, err = pgx.CollectOneRow(rows, pgx.RowToStructByPos[Event])
	if err != nil {
		return e.event, err
	}

	return e.event, nil
}
