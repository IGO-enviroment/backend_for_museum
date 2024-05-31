package usecase

import (
	entity_client "museum/app/entity/client"
	repo_client "museum/app/repo/client"
)

type ContentSearchCase struct {
	repo   *repo_client.ContentSearchRepo
	entity *entity_client.SearchEntity
}

func NewContentSearch(
	repo *repo_client.ContentSearchRepo,
	entity *entity_client.SearchEntity,
) *ContentSearchCase {
	return &ContentSearchCase{
		repo:   repo,
		entity: entity,
	}
}

func (c *ContentSearchCase) Call() string {
	return "123"
}
