package admin

import (
	"context"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type CreateEventRepo struct {
	db      *postgres.Postgres
	l       *logger.Logger
	context context.Context
}

func NewCreateEventRepo(db *postgres.Postgres, l *logger.Logger) CreateEventRepo {
	return CreateEventRepo{
		db:      db,
		l:       l,
		context: context.Background(),
	}
}

func (c *CreateEventRepo) Call(event map[string]interface{}) (int, error) {
	var id int

	tx, err := c.db.Pool.Begin(c.context)
	if err != nil {
		return 0, err
	}

	event["created_at"] = time.Now()
	event["updated_at"] = time.Now()

	sql, data, err := c.db.Builder.Insert("events").SetMap(
		event,
	).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow(c.context, sql, data...).Scan(&id)
	if err != nil {
		return 0, err
	}

	if event["tag_ids"] != nil {
		err = c.createTagsLink(tx, id, event["tag_ids"].([]int))
		if err != nil {
			tx.Rollback(c.context)

			return 0, err
		}
	}

	err = tx.Commit(c.context)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (c *CreateEventRepo) ExistRecord(table string, ids []int) bool {
	var id int

	sql, args, err := c.db.Builder.Select("id").From(table).Where(
		squirrel.Eq{"id": ids},
	).ToSql()
	if err != nil {
		return false
	}

	err = c.db.Pool.QueryRow(c.context, sql, args...).Scan(&id)

	return err == nil
}

func (c *CreateEventRepo) createTagsLink(tx pgx.Tx, eventID int, tags []int) error {
	for tagID := range tags {
		sql, args, err := c.db.Builder.Insert("event_tags").Columns(
			"event_id", "tag_id", "created_at", "updated_at",
		).Values(
			eventID, tagID, time.Now(), time.Now(),
		).ToSql()
		if err != nil {
			return err
		}

		_, err = tx.Exec(c.context, sql, args...)
		if err != nil {
			return err
		}
	}

	return nil
}
