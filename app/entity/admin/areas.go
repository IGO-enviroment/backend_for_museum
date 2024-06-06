package entity

import "time"

// Выдача площадок для выбора.
type AreaIDSEntity struct {
	ID    int
	Title string
}

type CreateAreaEntity struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Publish      bool    `json:"publish"`
	AddressValue string  `json:"addressValue"`
	Longitude    float64 `json:"longitude" validate:"required"`
	Latitude     float64 `json:"latitude"  validate:"required"`
}

type GetAreaEntity struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Publish      bool      `json:"publish"`
	AddressValue string    `json:"addressValue"`
	CreatedAt    time.Time `json:"createdAt"`
	Longitude    float64   `json:"longitude"`
	Latitude     float64   `json:"latitude"`
}
