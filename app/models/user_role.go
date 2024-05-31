package models

import "time"

// Таблица храненеия ролей пользователя.
type UserRole struct {
	ID        int
	UserID    string
	RoleID    string
	CreatedAt *time.Time
}
