package entity

import (
	"mime/multipart"
	"time"
)

// Создание мероприятия.
type CreateEventEntity struct {
	Title           string
	Description     *string
	Duration        *int
	StartAt         *string
	TicketCount     *int
	Area            *int
	Type            *int
	Tags            *[]int
	PreviewImage    *multipart.FileHeader
	AdditionalFiles []*multipart.FileHeader
}

type ShowEventEntity struct {
	Tags []ShowEventTags
}
type ShowEventTags struct {
	ID   int
	Name string
}

// Публикация мероприятия.
type PublishEventEntity struct {
	ID int
}

// Все мероприятия для событий.
type EventTable struct {
	Events []EventForTable `json:"events"`
}

// Элементы в таблицe.
type EventForTable struct {
	ID          int        `json:"id"`
	Publish     bool       `json:"publish"`
	Title       string     `json:"title"`
	TicketCount int        `json:"ticketCount"`
	Type        *string    `json:"type"`
	Area        *string    `json:"area"`
	StartAt     *time.Time `json:"start_at"`
	CreatedAt   time.Time  `json:"created_at"`
}