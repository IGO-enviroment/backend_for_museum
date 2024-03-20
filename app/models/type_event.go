package models

import "time"

// Все типы событий в музее.
type TypeEvent struct {
	ID        int
	Named     string
	Propery   string
	Publish   bool
	CreatedAt *time.Time
}
