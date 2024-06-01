package entity

import (
	"time"

	"github.com/jackc/pgx/v5"
)

// Создание мероприятия.
type CreateEventEntity struct {
	Title       string
	Description string
	StartAt     time.Time
	TicketCount int
	Area        int
	Type        int
	Tags        []int
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
	ID          int       `json:"id"`
	Publish     bool      `json:"publich"`
	Title       string    `json:"title"`
	TicketCount int       `json:"ticketCount"`
	Type        string    `json:"type"`
	Area        string    `json:"area"`
	CreatedAt   time.Time `json:"created_at"`
}

func (e *EventTable) ScanFromEquery(rows pgx.Rows) (EventTable, error) {
	result := EventTable{
		Events: []EventForTable{},
	}

	for rows.Next() {
		var item EventForTable

		err := rows.Scan(
			&item.ID,
			&item.Title,
			&item.Publish,
			&item.TicketCount,
			&item.Type,
			&item.Area,
			&item.CreatedAt,
		)
		if err != nil {
			return result, err
		}

		result.Events = append(result.Events, item)
	}

	return result, nil
}
