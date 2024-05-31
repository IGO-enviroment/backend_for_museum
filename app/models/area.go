package models

import "time"

// Площадки музея.
type Area struct {
	ID          int
	Name        string
	Description string
	Publish     bool
	Address     string
	CreatedAt   *time.Time
}
