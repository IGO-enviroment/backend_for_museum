package repo

import (
	"context"
	"github.com/Masterminds/squirrel"
	entity_test "museum/app/entity/test"
	"museum/app/models"
	"museum/pkg/logger"
	"museum/pkg/postgres"
	"time"
)

const (
	tableName = "contents"
)

type TextEditorRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewTextEditorRepo(db *postgres.Postgres, l *logger.Logger) *TextEditorRepo {
	return &TextEditorRepo{
		db: db,
		l:  l,
	}
}

func (t TextEditorRepo) CreateTextContent(textContent *entity_test.TestTextEditorContent) (int, error) {
	content := mapToDbContent(textContent)
	sql, args, err := t.db.Builder.Insert(tableName).
		Columns("type_value", "data_value", "order_value", "model_id", "model_type", "options", "created_at").
		Values(content.TypeValue, content.DataValue, content.OrderValue, content.ModelID, content.ModelType, content.Options, content.CreatedAt).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		t.l.Error("Unable to build INSERT query", err)
		return 0, err
	}
	var id int
	rows, err := t.db.Pool.Query(context.Background(), sql, args...)
	if err != nil {
		t.l.Error("Unable to query insert", err)
		return 0, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			t.l.Error("Unable to scan INSERT query", err)
			return 0, err
		}
		return id, nil
	}
	return 0, nil
}

func (t TextEditorRepo) GetTextContent(contentId int) (*entity_test.TestTextEditorContent, error) {
	query, args, err := t.db.Builder.
		Select(t.selectFields()...).
		From(tableName).
		Where(squirrel.Eq{"id": contentId}).
		ToSql()
	if err != nil {
		t.l.Error("Unable to build SELECT query", err)
		return nil, err
	}
	var contentEntity entity_test.TestTextEditorContent
	rows, err := t.db.Pool.Query(context.Background(), query, args...)
	if err != nil {
		t.l.Error("Unable to query insert", err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(
			&contentEntity.Id,
			&contentEntity.TypeValue,
			&contentEntity.DataValue,
			&contentEntity.OrderValue,
			&contentEntity.ModelID,
			&contentEntity.ModelType,
			&contentEntity.Options,
		)
		if err != nil {
			t.l.Error("Unable to scan query", err)
			return nil, err
		}
		return &contentEntity, nil
	}
	return nil, nil
}

func (t TextEditorRepo) UpdateTextContent(contentId int, newText string) error {
	query, args, err := t.db.Builder.
		Update(tableName).
		Set("data_value", newText).
		Where(squirrel.Eq{"id": contentId}).
		ToSql()
	if err != nil {
		t.l.Error("Unable to build UPDATE query", err)
		return err
	}
	_, err = t.db.Pool.Exec(context.Background(), query, args...)
	if err != nil {
		t.l.Error("Unable to query update", err)
		return err
	}
	return nil
}

func mapToDbContent(textContent *entity_test.TestTextEditorContent) models.Content {
	t := time.Now()
	return models.Content{
		TypeValue:  textContent.TypeValue,
		DataValue:  textContent.DataValue,
		OrderValue: textContent.OrderValue,
		ModelID:    textContent.ModelID,
		ModelType:  textContent.ModelType,
		Options:    textContent.Options,
		CreatedAt:  &t,
	}
}

func (t TextEditorRepo) selectFields() []string {
	return []string{
		"id", "type_value",
		"data_value", "order_value",
		"model_id", "model_type",
		"options",
	}
}
