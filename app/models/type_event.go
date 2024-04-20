package models

import "time"

// Все типы событий в музее.
type TypeEvent struct {
	ID          int
	Name        string
	Description string
	Publish     bool
	CreatedAt   *time.Time
}
