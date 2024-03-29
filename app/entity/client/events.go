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
