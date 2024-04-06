package usecase

import (
	entity_test "museum/app/entity/test"
	repo_test "museum/app/repo/test"
)

type TestTextEditorGetCase struct {
	repo      *repo_test.TextEditorRepo
	contentId int
}

func NewTestTextEditorGetCase(repo *repo_test.TextEditorRepo, contentId int) *TestTextEditorGetCase {
	return &TestTextEditorGetCase{
		repo:      repo,
		contentId: contentId,
	}
}

func (t TestTextEditorGetCase) Call() (*entity_test.TestTextEditorContent, error) {
	return t.repo.GetTextContent(t.contentId)
}
