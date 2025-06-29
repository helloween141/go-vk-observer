package vk

import (
	"database/sql"
	"go-vk-observer/internal/pkg/interfaces"
	"go-vk-observer/internal/pkg/utils"
	"log"
)

type Handler struct {
	vkClient           Client
	telegramSender     interfaces.TelegramSenderInterface
	telegramRepository interfaces.TelegramRepositoryInterface
}

func NewHandler(vkClient Client, telegramSender interfaces.TelegramSenderInterface, telegramRepository interfaces.TelegramRepositoryInterface) *Handler {
	return &Handler{vkClient: vkClient, telegramSender: telegramSender, telegramRepository: telegramRepository}
}

func (handler *Handler) HandleNotifications() error {
	telegramNotifications, err := handler.telegramRepository.GetList()
	if err != nil {
		return err
	}
	for _, telegramNotification := range telegramNotifications {
		response, err := handler.vkClient.SendGetWallRequest(telegramNotification.Slug)
		if err != nil {
			continue
		}

		items := response.Response.Items
		for i := len(items) - 1; i >= 0; i-- {
			if items[i].Date <= telegramNotification.LastPostDate.Int64 {
				continue
			}

			err := handler.telegramSender.SendVkPost(
				telegramNotification.TelegramID,
				telegramNotification.Slug,
				telegramNotification.Name,
				utils.FormatTimestampToDatetime(items[i].Date),
				items[i].Text,
			)
			if err != nil {
				log.Println(err)
				continue
			}
		}

		err = handler.telegramRepository.Update(
			telegramNotification.TelegramID,
			telegramNotification.EntityID,
			sql.NullInt64{Int64: items[0].Date, Valid: true},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
