package models

import (
	"museum/pkg/postgres"
	"time"
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

	CreatedAt time.Time
	UpdateAt  time.Time
}

type EventModel struct {
	db *postgres.Postgres
}
