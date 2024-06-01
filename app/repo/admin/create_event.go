package admin

import (
	"context"
	"museum/app/models"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"
)

type CreateEventRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewCreateEventRepo(db *postgres.Postgres, l *logger.Logger) CreateEventRepo {
	return CreateEventRepo{
		db: db,
		l:  l,
	}
}

func (c *CreateEventRepo) Call(event *models.Event) (int, error) {
	var id int

	event.CreatedAt, event.UpdateAt = time.Now(), time.Now()

	sql, data, err := c.db.Builder.Insert("events").Columns(
		"title", "start_at", "created_at", "updated_at",
	).Values(
		event.Title, event.StartAt, event.CreatedAt, event.UpdateAt,
	).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		return 0, err
	}

	err = c.db.Pool.QueryRow(context.Background(), sql, data...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
