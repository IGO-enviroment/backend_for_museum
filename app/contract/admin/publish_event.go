package admin

type PublishEvent struct {
	ID int `json:"id" validate:"required,gte=0"`
}
