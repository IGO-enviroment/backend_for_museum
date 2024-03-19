package tasks

import (
	"museum/config"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

// Крон задачи.
type Tasks struct {
	cfg *config.Config
	db  *postgres.Postgres
	l   *logger.Logger
}

func NewTasks(cfg *config.Config) *Tasks {
	return &Tasks{
		cfg: cfg,
	}
}

func (t *Tasks) Run() {

}

func (t *Tasks) dependency() {

}
