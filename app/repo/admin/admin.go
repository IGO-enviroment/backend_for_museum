package admin

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	entity "museum/app/entity/admin"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

const (
	usersTable = "users"
)

type Repo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewAdminRepo(db *postgres.Postgres, l *logger.Logger) Repo {
	return Repo{
		db: db,
		l:  l,
	}
}

func (r Repo) GetAdmin(email string) (*entity.Admin, error) {
	admin := &entity.Admin{}
	getUserSql, args, err := r.db.Builder.
		Select(
			"u.id",
			"u.is_admin",
			"u.password_digest",
			"r.named").
		From(fmt.Sprintf("%s as u", usersTable)).
		InnerJoin("user_roles ur ON ur.user_id = u.id").
		InnerJoin("roles r ON r.id = ur.role_id").
		Where(squirrel.Eq{"email": email}).
		ToSql()

	if err != nil {
		r.l.Error(err)

		return nil, err
	}

	rows, err := r.db.Pool.Query(context.Background(), getUserSql, args...)
	if err != nil {
		r.l.Error(err)

		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&admin.Id, &admin.IsAdmin, &admin.Password, &admin.Role)
		if err != nil {
			r.l.Error(err)

			return nil, err
		}

		return admin, nil
	}

	return nil, nil
}
