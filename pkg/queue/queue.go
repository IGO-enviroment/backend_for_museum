package queue

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

// Queue -.
type Queue struct {
	Client *asynq.Client
}

var queueClient *Queue

// New -.
func New(url string, opts ...Option) (*Queue, error) {
	queueClient = &Queue{}

	// Custom options
	for _, opt := range opts {
		opt(queueClient)
	}

	queueClient.Client = asynq.NewClient(asynq.RedisClientOpt{Addr: url})
	return queueClient, nil
}

func ConnQueue() *Queue {
	return queueClient
}

// Close -.
func (q *Queue) Close() {
	if q.Client != nil {
		q.Client.Close()
	}
}
