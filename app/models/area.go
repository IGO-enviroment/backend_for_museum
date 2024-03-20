package models

import "time"

// Площадки музея.
type Area struct {
	ID        int
	Named     string
	Propery   string
	Publish   bool
	CreatedAt *time.Time
}
