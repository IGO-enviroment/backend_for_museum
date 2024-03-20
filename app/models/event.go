package models

import "time"

// Мероприятия.
type Event struct {
	ID      int
	Title   string
	Publish bool

	TicketCount int

	StartAt *time.Time

	Duration int

	AreaID int
	TypeID int

	CreatedAt *time.Time
}
