package admin

import (
	"context"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
)

type CreateContentBlocksRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewContentBlocksRepo(db *postgres.Postgres, l *logger.Logger) *CreateContentBlocksRepo {
	return &CreateContentBlocksRepo{
		db: db,
		l:  l,
	}
}

func (c *CreateContentBlocksRepo) CreateBlock(
	parentID, index int, parentType, typeBlock, value string,
) bool {
	var id int

	sql, args, err := c.db.Builder.Insert("content").Columns(
		"model_id", "model_type", "type", "value", "order_value",
		"created_at", "updated_at",
	).Values(
		parentID, parentType, typeBlock, value, index,
		time.Now(), time.Now(),
	).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		return false
	}

	err = c.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)

	return err == nil
}

func (c *CreateContentBlocksRepo) ExistParentRecord(id int, table string) bool {
	var findID int

	sql, args, err := c.db.Builder.Select("id").From(table).Where(
		squirrel.Eq{"id": id},
	).Limit(1).ToSql()
	err = c.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&findID)

	return err == nil
}
