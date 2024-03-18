package seeds

import (
	"fmt"
	"log"
	"museum/config"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

type Seeds struct {
	cfg *config.Config
	db  *postgres.Postgres
	l   *logger.Logger
}

func NewSeeds(cfg *config.Config) *Seeds {
	return &Seeds{
		cfg: cfg,
	}
}

// Запуск создания сидов.
func (s *Seeds) Run() {
	s.initDependency()
}

func (s *Seeds) CreateTypeEvents() {
	for typeEvent := range TypeEventsNames() {
		fmt.Println(typeEvent)
	}
}

func (s *Seeds) CreateTags() {

}

func (s *Seeds) CreateAreas() {

}

func (s *Seeds) CreateEvents(id int) {

}

func (s *Seeds) CreateContent(id int) {

}

func (s *Seeds) CreateEventTags(id int) {

}

func (s *Seeds) CreateFilters() {

}

func (s *Seeds) CreateUsers() {

}

func (s *Seeds) CreateUserRoles() {

}

func (s *Seeds) CreateRoles() {

}

func (s *Seeds) initDependency() {
	var err error

	s.db, err = postgres.New(s.cfg.PG.URL)

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer s.db.Close()
}
