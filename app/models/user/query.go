package users

import (
	"context"
	"museum/pkg/postgres"

	"github.com/Masterminds/squirrel"
)

type UserQuery struct {
	TabelName string

	u *User

	db *postgres.Postgres
}

func NewQuery() *UserQuery {
	return &UserQuery{TabelName: "users", u: NewUser(), db: postgres.DB()}
}

func (q *UserQuery) FindByID(id int) (*User, error) {
	sql, args, err := q.Store().Where("id = ?", id).Limit(1).ToSql()
	if err != nil {
		return nil, err
	}
	if err := q.db.Pool.QueryRow(context.Background(), sql, args...).Scan(q.u.Attr.MapField(q.u)...); err != nil {
		return nil, err
	}
	return q.u, nil
}

func (q *UserQuery) IsAdmin() bool {
	var name string
	sql, args, err := q.db.Builder.Select("name").From("roles").Where("user_id = ? AND name = ?", q.u.ID, "admin").Limit(1).ToSql()
	if err != nil {
		return false
	}
	if err := q.db.Pool.QueryRow(context.Background(), sql, args...).Scan(&name); err != nil {
		return false
	}
	return name == "admin"
}

func (q *UserQuery) SelectMin() string {
	return q.u.Attr.SelectMin()
}

func (q *UserQuery) Store() squirrel.SelectBuilder {
	return q.db.Builder.Select(q.SelectMin()).From(q.TabelName)
}
