package verify

import (
	"context"
	"museum/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

type VerifyQuery struct {
	TabelName string

	v *Verify

	db *postgres.Postgres
}

func NewQuery() *VerifyQuery {
	return &VerifyQuery{TabelName: "verifies", v: NewVerify(), db: postgres.DB()}
}

func (q *VerifyQuery) FindByID(id int) (*Verify, error) {
	sql, args, err := q.Store().Where("id = ?", id).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}
	if err := q.db.Pool.QueryRow(context.Background(), sql, args...).Scan(q.v.Attr.MapField(q.v)...); err != nil {
		return nil, err
	}
	return q.v, nil
}

func (q *VerifyQuery) FindByCode(code string) (*Verify, error) {
	sql, args, err := q.Store().Where("uniq_code = ?", code).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}
	if err := q.db.Pool.QueryRow(context.Background(), sql, args...).Scan(q.v.Attr.MapField(q.v)...); err != nil {
		return nil, err
	}
	return q.v, nil
}

func (q *VerifyQuery) SelectMin() string {
	return q.v.Attr.SelectMin()
}

func (q *VerifyQuery) Store() squirrel.SelectBuilder {
	return q.db.Builder.Select(q.SelectMin()).From(q.TabelName)
}
