package models

import (
	"context"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// Статьи с информацией.
type Info struct {
	ID          int
	Title       string
	Description string

	IsVisible bool

	PublishedAt time.Time

	PreviewURL string

	CreatedAt time.Time
	UpdateAt  time.Time
}

type InfoModel struct {
	info Info
	db   *postgres.Postgres
}

func NewInfoModel(db *postgres.Postgres) *InfoModel {
	return &InfoModel{
		db: db,
	}
}

// Поиск статьи по ID.
func (c *InfoModel) Find(id int) (Info, error) {
	sql, args, err := c.db.Builder.Select("*").From(
		"informations",
	).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return c.info, err
	}

	rows, err := c.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		return c.info, err
	}

	c.info, err = pgx.CollectOneRow(rows, pgx.RowToStructByPos[Info])
	if err != nil {
		return c.info, err
	}

	return c.info, nil
}
