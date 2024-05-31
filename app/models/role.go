package models

import "time"

// Роли доступные пользователям.
type Role struct {
	ID        int
	Named     string
	Propery   string
	CreatedAt *time.Time
}
