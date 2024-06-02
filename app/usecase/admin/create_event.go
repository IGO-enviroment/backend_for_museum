package admin

import (
	"museum/app/api/external"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
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
	var previewURL string
	var err error
	var ok bool

	if e.entity.PreviewImage != nil {
		previewURL, ok = e.uploadPreview()
		if !ok {
			return 0, &e.errorResp
		}
	}

	id, err := e.repo.Call(e.CollectData(&previewURL))
	if err != nil {
		e.errorResp.Msg = "Неизвестная ошибка"

		return 0, &e.errorResp
	}

	return id, nil
}

func (e *CreateEventCase) CollectData(previewURL *string) map[string]interface{} {
	data := map[string]interface{}{
		"title": e.entity.Title,
	}

	if e.entity.Description != nil && *e.entity.Description != "" {
		data["description"] = *e.entity.Description
	}

	if e.entity.StartAt != nil && !(*e.entity.StartAt).IsZero() {
		data["start_at"] = *e.entity.StartAt
	}

	if e.entity.Area != nil && *e.entity.Area != 0 && e.repo.ExistRecord("areas", []int{*e.entity.Area}) {
		data["area_id"] = *e.entity.Area
	}

	if e.entity.Type != nil && *e.entity.Type != 0 && e.repo.ExistRecord("type_events", []int{*e.entity.Type}) {
		data["type_id"] = *e.entity.Type
	}

	if e.entity.Tags != nil && len(*e.entity.Tags) != 0 && e.repo.ExistRecord("tags", *e.entity.Tags) {
		data["tag_ids"] = *e.entity.Tags
	}

	if e.entity.TicketCount != nil && *e.entity.TicketCount != 0 {
		data["ticket_count"] = *e.entity.TicketCount
	}

	if e.entity.Duration != nil && *e.entity.Duration != 0 {
		data["duration"] = *e.entity.Duration
	}

	if previewURL != nil && *previewURL != "" {
		data["preview_url"] = *previewURL
	}

	return data
}

func (e *CreateEventCase) uploadPreview() (string, bool) {
	objectURL, err := external.NewUploadFileAPI(e.entity.PreviewImage).UploadObject()
	if err != nil {
		e.errorResp.Errors = append(e.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "previewFile",
				Value: "Ошибка при сохранении изображения",
			},
		)

		return "", false
	}

	return objectURL, true
}
