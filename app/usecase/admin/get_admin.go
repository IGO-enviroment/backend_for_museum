package admin

import (
	entity "museum/app/entity/admin"
	admin_repo "museum/app/repo/admin"
)

type GetAdminCase struct {
	repo  admin_repo.Repo
	email string
}

func NewGetAdminCase(
	repo admin_repo.Repo,
	email string) GetAdminCase {
	return GetAdminCase{
		repo,
		email,
	}
}

func (a *GetAdminCase) Call() (*entity.Admin, error) {
	return a.repo.GetAdmin(a.email)
}
