package models

import "time"

// Общая модель пользователя.
type User struct {
	ID             int
	Email          string
	DigestPassword string
	IsAdmin        bool
	CreatedAt      *time.Time
}
