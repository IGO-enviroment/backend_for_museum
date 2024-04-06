package usecase

import (
	repo_test "museum/app/repo/test"
)

type TestTextEditorUpdateCase struct {
	repo      *repo_test.TextEditorRepo
	contentId int
	newText   string
}

func NewTestTextEditorUpdateCase(
	repo *repo_test.TextEditorRepo,
	contentId int,
	newText string) *TestTextEditorUpdateCase {
	return &TestTextEditorUpdateCase{
		repo:      repo,
		contentId: contentId,
		newText:   newText,
	}
}

func (t TestTextEditorUpdateCase) Call() error {
	return t.repo.UpdateTextContent(t.contentId, t.newText)
}
