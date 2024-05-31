package models

import "time"

// Теги/категории.
type Tag struct {
	ID          int
	Name        string
	Description string
	Group       string
	CreatedAt   *time.Time
}

func (t *Tag) Groups() map[string]string {
	return map[string]string{
		"ByAge": "byage",
		"Else":  "else",
	}
}
