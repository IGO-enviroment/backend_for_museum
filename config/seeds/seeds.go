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
	cfg          *config.Config
	db           *postgres.Postgres
	typeIds      []int
	tagIds       []int
	areaIds      []int
	eventsIds    []int
	eventTypeIds []int
}

func NewSeeds(cfg *config.Config) *Seeds {
	return &Seeds{
		cfg: cfg,
	}
}

// Запуск создания сидов.
func (s *Seeds) Run() {
	s.initDependency()
	s.CreateTags()
	s.CreateAreas()
	s.CreateTypeEvents()
	s.CreateEvents()
	s.CreateEventTags()
}

func (s *Seeds) CreateTypeEvents() {
	var id int
	var sql string
	var args []interface{}
	var err error

	for _, data := range TypeEventsData() {
		sql, args, err = s.db.Builder.Insert("type_events").Columns(
			"name", "description", "publish", "created_at", "updated_at",
		).Values(
			data...,
		).Suffix("RETURNING \"id\"").ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.typeIds = append(s.typeIds, id)
	}
}

func (s *Seeds) CreateTags() {
	var id int
	var sql string
	var args []interface{}
	var err error

	for _, data := range TagsData() {
		sql, args, err = s.db.Builder.Insert("tags").Columns(
			"name", "description", "group_name", "created_at", "updated_at",
		).Values(
			data...,
		).Suffix("RETURNING \"id\"").ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.tagIds = append(s.tagIds, id)
	}
}

func (s *Seeds) CreateAreas() {
	var id int
	var sql string
	var args []interface{}
	var err error

	for _, data := range AreasData() {
		sql, args, err = s.db.Builder.Insert("areas").Columns(
			"name", "description", "publish", "address_value", "created_at", "updated_at",
		).Values(
			data...,
		).Suffix("RETURNING \"id\"").ToSql()
		if err != nil {
			log.Fatalf("%s", err)
		}

		err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.areaIds = append(s.areaIds, id)
	}
}

func (s *Seeds) CreateEvents() {
	var id int
	var sql string
	var args []interface{}
	var err error

	for _, eventData := range EventsData() {
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
			s.randId(s.areaIds),
			s.randId(s.typeIds),
		).Suffix("RETURNING \"id\"").ToSql()

		if err != nil {
			log.Fatalf("%s", err)
		}

		err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.eventsIds = append(s.eventsIds, id)
	}
}

func (s *Seeds) CreateEventTags() {
	var id int
	var sql string
	var args []interface{}
	var err error

	for _, eventId := range s.eventsIds {
		sql, args, err = s.db.Builder.Insert("event_tags").Columns(
			"event_id", "tag_id",
		).Values(
			eventId,
			s.randId(s.tagIds),
		).Suffix("RETURNING \"id\"").ToSql()

		if err != nil {
			log.Fatalf("%s", err)
		}

		err = s.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&id)
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.eventTypeIds = append(s.eventTypeIds, id)
	}
}

func (s *Seeds) CreateContent(id int) {

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

func (s *Seeds) initDependency() {
	var err error

	s.db, err = postgres.New(s.cfg.PG.URL)

	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
}

func (s *Seeds) randId(data []int) int {
	return data[rand.Intn(len(data))]
}
