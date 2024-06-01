package admin

import (
	"museum/pkg/logger"
	"museum/pkg/postgres"
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

func (c *CreateContentBlocksRepo) Create() {
}
