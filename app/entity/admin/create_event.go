package entity

import "time"

type CreateEventEntity struct {
	Title       string
	Description string
	StartAt     time.Time
	TicketCount int
	Area        int
	Type        int
	Tags        []int
}
