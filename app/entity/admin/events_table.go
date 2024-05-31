package entity

import (
	"time"

	"github.com/jackc/pgx/v4"
)

type EventForTable struct {
	ID          int       `json:"id" db:"id"`
	Publish     bool      `json:"publich" db:"publish"`
	Title       string    `json:"title" db:"title"`
	TicketCount int       `json:"ticketCount" db:"ticket_count"`
	Type        string    `json:"type" db:"type"`
	Area        string    `json:"area" db:"area"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type EventTable struct {
	Events []EventForTable `json:"events"`
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
