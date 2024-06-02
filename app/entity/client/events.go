package entity

import "time"

type TextForEvent struct {
	Data string
	By   string
}

type DateForEvent struct {
	Start time.Time
	End   time.Time
}

type EventsEntity struct {
	Text        TextForEvent
	Date        DateForEvent
	Tags        []int
	Areas       []int
	Types       []int
	Duration    int
	TypeTicket  []string
	TicketCount []string
	Price       []string
	Page        *int `json:"page,omitempty"`
}

type SelectValue struct {
	ID         int    `json:"id"`
	IsSelected bool   `json:"isSelected"`
	Name       string `json:"name"`
}

type RangeValue struct {
	Max          int `json:"max"`
	Min          int `json:"min"`
	SelectedTo   int `json:"to"`
	SelectedFrom int `json:"from"`
}

type FilterItem struct {
	Name   string        `json:"name"`
	Title  string        `json:"title"`
	Type   string        `json:"type"`
	Values []interface{} `json:"values"`
}

type EventItem struct {
	Name         string   `json:"name" db:"name"`
	Date         string   `json:"date" db:"start_at"`
	Area         string   `json:"area" db:"area_id"`
	Type         string   `json:"type" db:"type_id"`
	Duration     string   `json:"duration" db:"duration"`
	TicketsCount int      `json:"ticketsCount"`
	PreviewUrl   string   `json:"previewUrl" db:"preview_url"`
	Price        string   `json:"price"`
	Tags         []string `json:"tags"`
}

type PagePagination struct {
	Total   int `json:"total"`
	Current int `json:"current"`
}

// Ответ по страницы фильтрации.
type EventsResponse struct {
	Filters []FilterItem   `json:"filters"`
	Events  []EventItem    `json:"events"`
	Page    PagePagination `json:"page"`
}

type ShowEvent struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PreviewURL  string `json:"previewUrl"`
	Images      string `json:"images"`
}
