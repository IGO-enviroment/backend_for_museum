// Package postgres implements postgres connection.
package postgres

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

var (
	_globalPG *Postgres
	_oncePG   sync.Once
)

// New -.
func New(url string, opts ...Option) (*Postgres, error) {
	var err error

	_oncePG.Do(func() {
		var poolConfig *pgxpool.Config

		_globalPG = &Postgres{
			maxPoolSize:  _defaultMaxPoolSize,
			connAttempts: _defaultConnAttempts,
			connTimeout:  _defaultConnTimeout,
		}

		// Custom options
		for _, opt := range opts {
			opt(_globalPG)
		}

		_globalPG.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		poolConfig, err = pgxpool.ParseConfig(url)
		if err != nil {
			return
		}

		poolConfig.MaxConns = int32(_globalPG.maxPoolSize)

		for _globalPG.connAttempts > 0 {
			_globalPG.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
			if err == nil {
				break
			}

			log.Printf("Postgres is trying to connect, attempts left: %d", _globalPG.connAttempts)

			time.Sleep(_globalPG.connTimeout)

			_globalPG.connAttempts--
		}
	})

	if err != nil {
		return nil, fmt.Errorf("postgres - NewPostgre: %w", err)
	}

	return _globalPG, nil
}

func DB() *Postgres {
	return _globalPG
}

// Close -.
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
