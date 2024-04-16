package seeds

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"museum/config"
	"museum/pkg/postgres"
)

type Seeds struct {
	cfg     *config.Config
	db      *postgres.Postgres
	typeIds []int
	tagIds  []int
	areaIds []int
}

func NewSeeds(cfg *config.Config) *Seeds {
	return &Seeds{
		cfg: cfg,
	}
}

// Запуск создания сидов.
func (s *Seeds) Run() {
	s.initDependency()
	// s.CreateTypeEvents()
	// s.CreateTags()
	// s.CreateAreas()
	s.CreateEvents()
}

func (s *Seeds) CreateTypeEvents() {
	var sql string
	var args []interface{}
	var err error

	for _, typeEvent := range TypeEventsNames() {
		sql, args, err = s.db.Builder.Insert("type_events").Columns("named", "property", "publish").Values(
			typeEvent, RandStringBytes(200), true,
		).ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		_, err = s.db.Pool.Exec(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
}

func (s *Seeds) CreateTags() {
	var sql string
	var args []interface{}
	var err error

	for _, tag := range TagsNames() {
		sql, args, err = s.db.Builder.Insert("tags").Columns("named", "property").Values(
			tag, RandStringBytes(200),
		).ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		_, err = s.db.Pool.Exec(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
}

func (s *Seeds) CreateAreas() {
	var sql string
	var args []interface{}
	var err error

	for _, tag := range AreasNames() {
		sql, args, err = s.db.Builder.Insert("areas").Columns(
			"named", "property", "publish", "address_value",
		).Values(
			tag, RandStringBytes(200), true, "address_value",
		).ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		_, err = s.db.Pool.Exec(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
}

func (s *Seeds) CreateEvents() {
	var sql string
	var args []interface{}
	var err error

	for _, eventData := range EventsData() {
		areaID := s.LastID("areas")
		typeID := s.LastID("type_events")

		sql, args, err = s.db.Builder.Insert("events").Columns(
			"title", "publish", "ticket_count",
			"start_at", "duration", "price", "area_id", "type_id",
		).Values(
			eventData.Name,
			eventData.Publish,
			eventData.TicketCount,
			eventData.StartAt,
			eventData.Duration,
			eventData.Price,
			rand.Intn(areaID-1)+1,
			rand.Intn(typeID-1)+1,
		).ToSql()

		if err != nil {
			log.Fatalf("%s", err)
		}

		_, err = s.db.Pool.Exec(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
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
	var sql string
	var args []interface{}
	var err error

	for _, role := range RolesNames() {
		sql, args, err = s.db.Builder.Insert("roles").Columns(
			"named", "property",
		).Values(
			role, RandStringBytes(200),
		).ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		_, err = s.db.Pool.Exec(context.Background(), sql, args...)
		if err != nil {
			log.Fatalf("%s", err)
		}
	}
}

func (s *Seeds) LastID(tableName string) int {
	var id int

	sql, args, err := s.db.Builder.Select("MAX(id)").From(tableName).ToSql()
	if err != nil {
		log.Fatalf("%s", err)
	}

	err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(id)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return id
}

func (s *Seeds) initDependency() {
	var err error

	s.db, err = postgres.New(s.cfg.PG.URL)

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
}
