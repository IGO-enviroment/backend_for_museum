package models

import "time"

// Сохраненные фильтрации для дальнейшего переиспользования.
type Filter struct {
	ID         int
	Named      string
	Property   string
	Publish    bool
	OrderValue int
	UserID     int
	Options    string
	CreatedAt  *time.Time
}
