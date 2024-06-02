package admin

import (
	"museum/app/api/external"
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	repo_admin "museum/app/repo/admin"
	"time"
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
	var additionalUrls []string
	var err error
	var ok bool

	if e.entity.PreviewImage != nil {
		previewURL, ok = e.uploadPreview()
		if !ok {
			return 0, &e.errorResp
		}
	}

	if len(e.entity.AdditionalFiles) != 0 {
		additionalUrls, ok = e.uploadAdditional()
		if !ok {
			return 0, &e.errorResp
		}
	}

	id, err := e.repo.Call(e.CollectData(&previewURL), &additionalUrls)
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

	if valid, startAt := e.isValidDate(e.entity.StartAt); valid {
		data["start_at"] = startAt
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

func (e *CreateEventCase) isValidDate(date *string) (bool, *time.Time) {
	if date != nil && *date != "" {
		return false, nil
	}

	if formattedDate, err := time.Parse(time.RFC3339, *date); err != nil {
		return true, &formattedDate
	}

	return false, nil
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

func (e *CreateEventCase) uploadAdditional() ([]string, bool) {
	var urls []string

	for _, addtionalFile := range e.entity.AdditionalFiles {
		objectURL, err := external.NewUploadFileAPI(addtionalFile).UploadObject()
		if err != nil {
			e.errorResp.Errors = append(e.errorResp.Errors,
				handlers.ErrorWithKey{
					Key:   "AdditionalFiles",
					Value: "Ошибка при сохранении изображений",
				},
			)

			return urls, false
		}

		urls = append(urls, objectURL)
	}

	return urls, true
}
