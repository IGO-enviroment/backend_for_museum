package entity

import "time"

// Выдача тегов для выбора.
type TagsIDSEntity struct {
	ID    int
	Title string
}

type TagEntity struct {
	Id          int
	Name        string
	Description string
	GroupName   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CreateTagEntity struct {
	Name        string
	Description string
	GroupName   string
}
