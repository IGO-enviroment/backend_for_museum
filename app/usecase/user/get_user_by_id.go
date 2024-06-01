package user

import (
	entity "museum/app/entity/user"
	user_repo "museum/app/repo/user"
)

type GetUserCase struct {
	repo user_repo.GetUserRepo
	id   int
}

func NewGetUserCase(
	repo user_repo.GetUserRepo,
	id int) GetUserCase {
	return GetUserCase{
		repo,
		id,
	}
}

func (c *GetUserCase) Call() (*entity.User, error) {
	return c.repo.GetUser(c.id)
}

func (c *GetUserCase) IsUserSuperAdmin() (bool, error) {
	return c.repo.IsUserSuperAdmin(c.id)
}
