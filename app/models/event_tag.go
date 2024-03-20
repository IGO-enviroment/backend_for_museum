package models

import "time"

// Теги мероприятия.
type EventTag struct {
	ID        int
	EventID   int
	TagID     int
	CreatedAt *time.Time
}
