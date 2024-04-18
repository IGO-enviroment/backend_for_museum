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
	Text     TextForEvent
	Date     DateForEvent
	Tags     []int
	Areas    []int
	Types    []int
	Duration int
	Tickets  int
	Price    int
	Page     int
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
	Name         string   `json:"name"`
	Date         string   `json:"date"`
	Area         string   `json:"area"`
	Type         string   `json:"type"`
	Duration     string   `json:"duration"`
	TicketsCount int      `json:"ticketsCount"`
	PreviewUrl   string   `json:"previewUrl"`
	Price        string   `json:"price"`
	Tags         []string `json:"tags"`
}

type PagePagination struct {
	Total   string `json:"total"`
	Current string `json:"current"`
}

// Ответ по страницы фильтрации
type EventsResponse struct {
	Filters []FilterItem   `json:"filters"`
	Events  []EventItem    `json:"events"`
	Page    PagePagination `json:"page"`
}
