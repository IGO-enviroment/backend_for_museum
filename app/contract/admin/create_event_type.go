package admin

type CreateEventType struct {
	Name        string `json:"name" validate:"max=255"`
	Description string `json:"description" validate:"max=255"`
	IsVisible   bool   `json:"isVisible"`
}
