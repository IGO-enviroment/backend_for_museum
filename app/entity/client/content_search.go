package entity

import "time"

// Глобальный поиск по всему контенту постов.
type SearchEntity struct {
	TypeSearch string
	Target     string
}

// Результат поиска.
type ResultContentSearch struct {
	Events       []FindContent `json:"events"`
	News         []FindContent `json:"news"`
	Informations []FindContent `json:"informations"`
	Page         int           `json:"page"`
}

// Найденный блок контнета.
type FindContent struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Value     string     `json:"value"`
	CreatedAt *time.Time `json:"createdAt"`
}
