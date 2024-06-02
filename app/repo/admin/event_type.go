package admin

import (
	"context"
	entity_admin "museum/app/entity/admin"
	"museum/app/models"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"

	"github.com/Masterminds/squirrel"
)

const (
	tableName = "type_events"
)

type EventTypeRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewEventTypeRepo(db *postgres.Postgres, l *logger.Logger) EventTypeRepo {
	return EventTypeRepo{
		db: db,
		l:  l,
	}
}

func (e EventTypeRepo) Create(entity entity_admin.EventTypeEntity) (int, error) {
	eventType := mapToEventTypeDb(entity)
	sql, args, err := e.db.Builder.Insert(tableName).
		Columns("name", "description", "is_visible", "created_at", "updated_at").
		Values(eventType.Name, eventType.Description, eventType.IsVisible, eventType.CreatedAt, eventType.UpdatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		e.l.Error("Unable to build INSERT query", err)
		return 0, err
	}
	var id int
	rows, err := e.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		e.l.Error("Unable to query insert", err)
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			e.l.Error("Unable to scan INSERT query", err)
			return 0, err
		}
		return id, nil
	}
	return 0, nil
}

func (e EventTypeRepo) Get() ([]entity_admin.EventTypeEntity, error) {
	var result []entity_admin.EventTypeEntity
	sql, args, err := e.db.Builder.
		Select(
			"name",
			"description",
			"is_visible",
		).
		From(tableName).
		ToSql()
	if err != nil {
		e.l.Error("Unable to build SELECT query", err)
		return nil, err
	}

	rows, err := e.db.Pool.Query(context.Background(), sql, args...)

	defer rows.Close()

	for rows.Next() {
		var eventType entity_admin.EventTypeEntity
		err := rows.Scan(&eventType.Name, &eventType.Description, &eventType.IsVisible)
		if err != nil {
			e.l.Error("Unable to scan INSERT query", err)
			return nil, err
		}
		result = append(result, eventType)
	}
	return result, nil
}

func mapToEventTypeDb(entity entity_admin.EventTypeEntity) models.EventType {
	timeNow := time.Now()
	return models.EventType{
		Name:        entity.Name,
		Description: entity.Description,
		IsVisible:   entity.IsVisible,
		CreatedAt:   &timeNow,
		UpdatedAt:   &timeNow,
	}
}
