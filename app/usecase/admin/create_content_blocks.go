package admin

import (
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
)

type CreateContentBlocksCase struct {
	repo      *admin_repo.CreateContentBlocksRepo
	entity    *entity_admin.CreateContentBlocksEntity
	errorResp handlers.ErrorStruct
}

func NewCreateContentBlocksCaseCase(
	repo *admin_repo.CreateContentBlocksRepo, entity *entity_admin.CreateContentBlocksEntity) *CreateContentBlocksCase {
	return &CreateContentBlocksCase{
		repo:   repo,
		entity: entity,
	}
}

// Создание нового блока контента.
func (c *CreateContentBlocksCase) Call() (bool, *handlers.ErrorStruct) {
	return true, &c.errorResp
}

func (c *CreateContentBlocksCase) ForInfoCreate() {
}

// // Загрузка в хранилище S3.
// func (c *CreateContentBlocksCase) uploadFile() {
// }
