package user

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	entity "museum/app/entity/user"
	"museum/pkg/logger"
	"museum/pkg/postgres"
)

const (
	usersTable = "users"
)

type GetUserRepo struct {
	db *postgres.Postgres
	l  *logger.Logger
}

func NewGetUserRepo(db *postgres.Postgres, l *logger.Logger) GetUserRepo {
	return GetUserRepo{
		db: db,
		l:  l,
	}
}

func (r GetUserRepo) GetUser(id int) (*entity.User, error) {
	user := &entity.User{}
	getUserSql, args, err := r.db.Builder.
		Select(
			"u.id",
			"u.is_admin",
			"u.email",
			"u.password_digest",
			"u.created_at",
			"r.named").
		From(fmt.Sprintf("%s as u", usersTable)).
		LeftJoin("user_roles ur ON ur.user_id = u.id").
		LeftJoin("roles r ON r.id = ur.role_id").
		Where(squirrel.Eq{"u.id": id}).
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
		err := rows.Scan(
			&user.ID,
			&user.IsAdmin,
			&user.Email,
			&user.DigestPassword,
			&user.CreatedAt,
			&user.Role,
		)

		if err != nil {
			r.l.Error(err)

			return nil, err
		}

		return user, nil
	}

	return nil, nil
}

func (r GetUserRepo) IsUserSuperAdmin(id int) (bool, error) {
	getUserSql, args, err := r.db.Builder.
		Select("r.named").
		From(fmt.Sprintf("%s as u", usersTable)).
		InnerJoin("user_roles ur ON ur.user_id = u.id").
		InnerJoin("roles r ON r.id = ur.role_id").
		Where(squirrel.Eq{"u.id": id}).
		Where(squirrel.Eq{"r.named": "SuperAdmin"}).
		ToSql()

	if err != nil {
		r.l.Error(err)

		return false, err
	}

	rows, err := r.db.Pool.Query(context.Background(), getUserSql, args...)
	if err != nil {
		r.l.Error(err)

		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
