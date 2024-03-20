package models

import "time"

// Теги/категории.
type Tag struct {
	ID        int
	Named     string
	Propery   string
	CreatedAt *time.Time
}
