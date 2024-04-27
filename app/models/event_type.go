package models

import "time"

// Все типы событий в музее.
type EventType struct {
	ID          int
	Name        string
	Description string
	IsVisible   bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
