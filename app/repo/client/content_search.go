package repo

import (
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type ContentSearchRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewContentSearchRepo(db *postgres.Postgres, l *logger.Logger) *ContentSearchRepo {
	return &ContentSearchRepo{
		db: db,
		l:  l,
	}
}

// Поиск по всем типам контента.
func (c *ContentSearchRepo) AllPosts() {
	c.Events()
	c.News()
	c.Information()
}

// Поиск по мероприятим.
func (c *ContentSearchRepo) Events() {}

// Поиск по новостям.
func (c *ContentSearchRepo) News() {}

// Поиск по информационным постам.
func (c *ContentSearchRepo) Information() {}
