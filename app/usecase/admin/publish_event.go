package admin

import (
	entity_admin "museum/app/entity/admin"
	"museum/app/handlers"
	"museum/app/models"
	repo_admin "museum/app/repo/admin"
)

type PublishEventCase struct {
	repo      *repo_admin.PublishEventRepo
	entity    *entity_admin.PublishEventEntity
	errorResp handlers.ErrorStruct
}

func NewPublishEventCase(repo *repo_admin.PublishEventRepo, entity *entity_admin.PublishEventEntity) *PublishEventCase {
	return &PublishEventCase{
		repo:   repo,
		entity: entity,
	}
}

func (p *PublishEventCase) Call() (bool, *handlers.ErrorStruct) {
	event, err := p.setEvent()

	ok := p.isValid(&event)
	if !ok {
		return false, &p.errorResp
	}

	err = p.repo.Call(p.entity.ID)
	if err != nil {
		p.errorResp.Msg = "Не известная ошибка"

		return false, &p.errorResp
	}

	return true, &p.errorResp
}

// описание хотя бы какое-то
// привязаны билеты
// площадка
// тип меры
// дата начала
// длительность
// картинка обязательна
func (p *PublishEventCase) isValid(event *models.Event) bool {
	return p.fieldsExist(event)
}

func (p *PublishEventCase) fieldsExist(event *models.Event) bool {
	result := true

	if event.Title == "" {
		p.errorResp.Errors = append(p.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "title",
				Value: "Укажите название мероприятия",
			},
		)

		result = false
	}

	if event.Description == "" {
		p.errorResp.Errors = append(p.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "description",
				Value: "Укажите описание мероприятия",
			},
		)

		result = false
	}

	if event.StartAt.IsZero() {
		p.errorResp.Errors = append(p.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "description",
				Value: "Укажите дату начала мероприятия",
			},
		)

		result = false
	}

	if event.Duration == 0 {
		p.errorResp.Errors = append(p.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "description",
				Value: "Укажите длительность мероприятия",
			},
		)

		result = false
	}

	if event.PreviewURL == "" {
		p.errorResp.Errors = append(p.errorResp.Errors,
			handlers.ErrorWithKey{
				Key:   "description",
				Value: "Укажите основное изображение",
			},
		)

		result = false
	}

	return result
}

func (p *PublishEventCase) setEvent() (models.Event, error) {
	event, err := p.repo.FindEvent(p.entity.ID)
	if err != nil {
		p.errorResp.Msg = "Ошибка при поиске мероприятия"

		return event, err
	}

	return event, nil
}
