package usecase

import (
	entity_test "museum/app/entity/test"
	repo_test "museum/app/repo/test"
)

type TestTextEditorCreateCase struct {
	repo   *repo_test.TextEditorRepo
	entity *entity_test.TestTextEditorContent
}

func NewTestTextEditorCreateCase(repo *repo_test.TextEditorRepo, entity *entity_test.TestTextEditorContent) *TestTextEditorCreateCase {
	return &TestTextEditorCreateCase{
		repo:   repo,
		entity: entity,
	}
}

func (t TestTextEditorCreateCase) Call() (int, error) {
	return t.repo.CreateTextContent(t.entity)
}
