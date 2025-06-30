package telegram

import (
	"context"
	"database/sql"
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

func (repository *Repository) Create(telegramID int64, entityID int32) error {
	_, err := repository.query.CreateTelegramNotification(context.Background(), dbstore.CreateTelegramNotificationParams{
		TelegramID: telegramID,
		EntityID:   entityID,
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repository *Repository) Delete(telegramID int64, entityID int32) error {
	_, err := repository.query.DeleteTelegramNotification(context.Background(), dbstore.DeleteTelegramNotificationParams{
		TelegramID: telegramID,
		EntityID:   entityID,
	})
	if err != nil {
		return err
	}

	return nil
}

func (repository *Repository) GetByTelegramID(telegramID int64) ([]dbstore.GetTelegramNotificationsByTelegramIDRow, error) {
	entities, err := repository.query.GetTelegramNotificationsByTelegramID(context.Background(), telegramID)

	if err != nil {
		return nil, err
	}

	return entities, nil
}

func (repository *Repository) IsEntityExistsByTelegramID(telegramID int64, entityID int32) bool {
	existsResult, err := repository.query.IsTelegramNotificationExists(context.Background(), dbstore.IsTelegramNotificationExistsParams{
		TelegramID: telegramID,
		EntityID:   entityID,
	})

	if err != nil {
		return false
	}

	return existsResult
}

func (repository *Repository) GetList() ([]dbstore.GetTelegramNotificationListRow, error) {
	list, err := repository.query.GetTelegramNotificationList(context.Background())

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (repository *Repository) Update(telegramID int64, entityID int32, lastPostDate sql.NullInt64) error {
	_, err := repository.query.UpdateTelegramNotification(context.Background(), dbstore.UpdateTelegramNotificationParams{
		TelegramID:   telegramID,
		EntityID:     entityID,
		LastPostDate: lastPostDate,
	})

	if err != nil {
		return err
	}

	return nil
}
