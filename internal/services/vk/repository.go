package vk

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-vk-observer/internal/database/gen/dbstore"
	"log"
)

type Repository struct {
	query *dbstore.Queries
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{query: dbstore.New(db)}
}

func (repository *Repository) GetBySlug(slug string) (*dbstore.VkEntity, error) {
	vkEntity, err := repository.query.GetVkEntityBySlug(context.Background(), slug)
	if err != nil {
		return nil, err
	}

	return &vkEntity, nil
}

func (repository *Repository) Create(slug string, name string) (*dbstore.VkEntity, error) {
	vkEntity, err := repository.query.CreateVkEntity(context.Background(), dbstore.CreateVkEntityParams{
		Slug: slug,
		Name: name,
	})
	if err != nil {
		log.Println(err)
		return &dbstore.VkEntity{}, err
	}

	return &vkEntity, nil
}
