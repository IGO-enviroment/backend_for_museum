package admin

import "time"

type CreateEvent struct {
	Title       string `form:"title" validate:"required|max_len:255" message:"required:Обязательное поле"`
	Description string `form:"description,omitempty" validate:"max_len:10000" message:"max:Слишком большое поле"`
	StartAt     string `form:"startAt,omitempty" validate:"-"`
	AreaID      int    `form:"area,omitempty" validate:"-"`
	TypeID      int    `form:"type,omitempty" validate:"-"`
	TagIDS      []int  `form:"tags,omitempty" validate:"-"`
	Duration    int    `form:"duration,omitempty" validate:"int"`
	TicketCount int    `form:"ticketCount,omitempty" validate:"int"`
}

type ShowEvent struct {
	ID int `json:"event_id" validate:"required|int"`
}

type CreateEventType struct {
	Name        string `json:"name" validate:"max=255"`
	Description string `json:"description" validate:"max=255"`
	IsVisible   bool   `json:"isVisible"`
}

type PublishEvent struct {
	ID int `json:"id" validate:"required,gte=0"`
}

type EventTypeById struct {
	Id          int `json:"id" validate:"required,gte=0"`
	Name        string
	Description string
	IsVisible   bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
