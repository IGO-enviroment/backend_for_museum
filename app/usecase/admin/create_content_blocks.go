package admin

import (
	"museum/app/api/external"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	admin_repo "museum/app/repo/admin"
)

const (
	parentInfo = "info"
	textType   = "text"
	fileType   = "file"
	imageType  = "image"
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
	ok := true

	if !c.isValid() {
		return false, &c.errorResp
	}

	if c.entity.ParentType == parentInfo {
		ok = c.ForInfoCreate()
	}

	return ok, &c.errorResp
}

func (c *CreateContentBlocksCase) ForInfoCreate() bool {
	if c.entity.Type == fileType || c.entity.Type == imageType {
		objectURL, ok := c.uploadFile()
		if !ok {
			return false
		}

		return c.CreateBlock(objectURL)
	}

	return c.CreateBlock(c.entity.ValueStr)
}

func (c *CreateContentBlocksCase) CreateBlock(value string) bool {
	result := c.repo.CreateBlock(
		c.entity.ParentID, c.entity.Index, c.entity.ParentType,
		c.entity.Type, value,
	)
	if !result {
		c.errorResp.Msg = "Ошибка при создании"

		return false
	}

	return result
}

func (c *CreateContentBlocksCase) uploadFile() (string, bool) {
	objectURL, err := external.NewUploadFileAPI(c.entity.ValueFile).UploadObject()
	if err != nil {
		c.errorResp.Msg = "Не удалось сохранить файл"

		return "", false
	}

	return objectURL, true
}

func (c *CreateContentBlocksCase) isValid() bool {
	if !c.isValidParent() {
		return false
	}

	if !c.isValidFields() {
		return false
	}

	return true
}

func (c *CreateContentBlocksCase) isValidParent() bool {
	validParents := []string{parentInfo}

	if !c.includeIn(c.entity.ParentType, validParents) {
		c.errorResp.Errors = append(c.errorResp.Errors,
			handlers.ErrorWithKey{
				Key: "parentType", Value: "Неизвестный тип родительской записи",
			},
		)

		return false
	}

	if !c.repo.ExistParentRecord(c.entity.ParentID, c.entity.ParentType) {
		c.errorResp.Errors = append(c.errorResp.Errors,
			handlers.ErrorWithKey{Key: "parentID", Value: "Родительская запись не найдена"},
		)

		return false
	}

	return true
}

func (c *CreateContentBlocksCase) isValidFields() bool {
	validTypes := []string{textType, fileType, imageType}

	if !c.includeIn(c.entity.Type, validTypes) {
		c.errorResp.Errors = append(c.errorResp.Errors,
			handlers.ErrorWithKey{
				Key: "type", Value: "Неверный тип",
			},
		)

		return false
	}

	if c.entity.Type == textType && c.entity.ValueStr == "" {
		c.errorResp.Errors = append(c.errorResp.Errors,
			handlers.ErrorWithKey{
				Key: "valueStr", Value: "Пустой контент блока",
			},
		)
	}

	if c.entity.Type != textType && c.entity.ValueFile == nil {
		c.errorResp.Errors = append(c.errorResp.Errors,
			handlers.ErrorWithKey{
				Key: "valueFile", Value: "Пустой контент блока",
			},
		)
	}

	return true
}

func (c *CreateContentBlocksCase) includeIn(value string, list []string) bool {
	finded := 0

	for _, typeTable := range list {
		if value != typeTable {
			continue
		}
		finded++
	}

	return finded == 1
}
