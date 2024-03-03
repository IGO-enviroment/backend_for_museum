// Темы событий, новостей и т.п.
package models

type Topic struct {
	TabelName string

	ID          int
	NameKey     string
	Description string
	Active      bool
}

func NewTopic() *Topic {
	return &Topic{TabelName: "topics"}
}
