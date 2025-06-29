package vk

import (
	"context"
	"github.com/jmoiron/sqlx"
	dbstore2 "go-vk-observer/internal/database/gen/dbstore"
	"log"
)

type Repository struct {
	query *dbstore2.Queries
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{query: dbstore2.New(db)}
}

func (repository *Repository) GetBySlug(slug string) (*dbstore2.VkEntity, error) {
	vkEntity, err := repository.query.GetVkEntityBySlug(context.Background(), slug)
	if err != nil {
		return nil, err
	}

	return &vkEntity, nil
}

func (repository *Repository) Create(slug string, name string, entityType string) (*dbstore2.VkEntity, error) {
	vkEntity, err := repository.query.CreateVkEntity(context.Background(), dbstore2.CreateVkEntityParams{
		Slug: slug,
		Name: name,
		Type: dbstore2.EntityType(entityType),
	})
	if err != nil {
		log.Println(err)
		return &dbstore2.VkEntity{}, err
	}

	return &vkEntity, nil
}
