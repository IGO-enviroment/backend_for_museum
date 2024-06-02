package admin

import (
	"museum/app/api/external"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	"museum/app/models"
	repo_admin "museum/app/repo/admin"
)

type CreateEventCase struct {
	repo      *repo_admin.CreateEventRepo
	entity    *entity_admin.CreateEventEntity
	errorResp handlers.ErrorStruct
}

func NewCreateEventCase(repo *repo_admin.CreateEventRepo, entity *entity_admin.CreateEventEntity) *CreateEventCase {
	return &CreateEventCase{
		repo:   repo,
		entity: entity,
	}
}

func (e *CreateEventCase) Call() (int, *handlers.ErrorStruct) {
	var err error

	if !e.isValid() {
		return 0, &e.errorResp
	}

	id, err := e.repo.Call(e.CollectData())
	if err != nil {
		e.errorResp.Msg = "Неизвестная ошибка"

		return 0, &e.errorResp
	}

	return id, nil
}

func (e *CreateEventCase) CollectData() *models.Event {
	data := models.Event{
		Title:   e.entity.Title,
		StartAt: e.entity.StartAt,
	}

	return &data
}

func (e *CreateEventCase) isValid() bool {
	return true
}

func (e *CreateEventCase) uploadPreview() (string, bool) {
	objectURL, err := external.NewUploadFileAPI(e.entity.PreviewImage).UploadObject()
	if err != nil {
		e.errorResp.Errors = append(e.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "previewURL",
				Value: "Ошибка при сохранении изображения",
			},
		)

		return "", false
	}

	return objectURL, true
}
