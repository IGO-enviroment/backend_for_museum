package models

import "time"

// Контент внутри мероприятия, новости, инструкции и т.п.
type Content struct {
	ID         int
	TypeValue  string
	DataValue  string
	OrderValue int
	ModelID    int
	ModelType  string
	Options    string
	CreatedAt  *time.Time
}
