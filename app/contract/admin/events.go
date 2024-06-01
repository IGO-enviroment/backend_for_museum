package admin

type CreateEvent struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=5000"`
	StartAt     string `json:"startAt,omitempty"`
}

type CreateEventType struct {
	Name        string `json:"name" validate:"max=255"`
	Description string `json:"description" validate:"max=255"`
	IsVisible   bool   `json:"isVisible"`
}

type PublishEvent struct {
	ID int `json:"id" validate:"required,gte=0"`
}
