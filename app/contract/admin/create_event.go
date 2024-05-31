package admin

type CreateEvent struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"max=5000"`
	StartAt     string `json:"startAt,omitempty"`
}
