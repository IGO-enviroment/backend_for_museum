package admin

import (
	"context"
	"museum/app/models"
	"museum/pkg/logger"
	"museum/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

type PublishEventRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewPublishEventRepo(db *postgres.Postgres, l *logger.Logger) *PublishEventRepo {
	return &PublishEventRepo{
		db: db,
		l:  l,
	}
}

func (p *PublishEventRepo) Call(id int) error {
	sql, data, err := p.db.Builder.Update("events").Set(
		"publish", true,
	).Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Pool.Exec(context.Background(), sql, data...)
	if err != nil {
		return err
	}

	return nil
}

func (p *PublishEventRepo) FindEvent(id int) (models.Event, error) {
	return models.NewEventModel(p.db).Find(id)
}
