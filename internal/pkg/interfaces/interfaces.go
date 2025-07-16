package interfaces

import (
	"database/sql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-vk-observer/internal/database/gen/dbstore"
	"go-vk-observer/internal/services/vk/responses"
)

type TelegramSenderInterface interface {
	GetBot() *tgbotapi.BotAPI
	SendMessage(telegramID int64, message string, isCommand bool) error
	SendMenu(telegramID int64) error
}

type TelegramRepositoryInterface interface {
	GetList() ([]dbstore.GetTelegramNotificationListRow, error)
	GetByTelegramID(telegramID int64) ([]dbstore.GetTelegramNotificationsByTelegramIDRow, error)
	IsEntityExistsByTelegramID(telegramID int64, entityID int32) bool
	Create(telegramID int64, entityID int32) error
	Delete(telegramID int64, entityID int32) error
	Update(telegramID int64, entityID int32, LstPostDate sql.NullInt64) error
}

type VkRepositoryInterface interface {
	GetBySlug(slug string) (*dbstore.VkEntity, error)
	Create(slug string, name string) (*dbstore.VkEntity, error)
}

type VkClientInterface interface {
	SendGetWallRequest(slug string) (*responses.WallResponse, error)
	SendGetGroupRequest(slug string) (*responses.GroupResponse, error)
	SendGetUserRequest(slug string) (*responses.UserResponse, error)
}
