package vk

import (
	"database/sql"
	"go-vk-observer/internal/pkg/interfaces"
	"log"
)

type Handler struct {
	vkClient           Client
	vkService          ServiceInterface
	telegramSender     interfaces.TelegramSenderInterface
	telegramRepository interfaces.TelegramRepositoryInterface
}

func NewHandler(vkClient Client, telegramSender interfaces.TelegramSenderInterface, telegramRepository interfaces.TelegramRepositoryInterface, vkService ServiceInterface) *Handler {
	return &Handler{vkClient: vkClient, telegramSender: telegramSender, telegramRepository: telegramRepository, vkService: vkService}
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

		lastPostDate := telegramNotification.LastPostDate.Int64
		posts := response.Response.Posts
		for i := len(posts) - 1; i >= 0; i-- {
			if posts[i].Date <= lastPostDate {
				continue
			}

			text, err := handler.vkService.CreatePostMessage(telegramNotification.Name, posts[i])
			if err != nil {
				log.Println("Can't create Telegram message", err)
				continue
			}

			err = handler.telegramSender.SendMessage(telegramNotification.TelegramID, text, false)
			if err != nil {
				log.Println("Can't send message to Telegram", err)
			}

			lastPostDate = posts[i].Date
		}

		err = handler.telegramRepository.Update(
			telegramNotification.TelegramID,
			telegramNotification.EntityID,
			sql.NullInt64{Int64: lastPostDate, Valid: true},
		)
		if err != nil {
			continue
		}
	}

	return nil
}
