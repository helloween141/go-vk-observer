package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-vk-observer/internal/pkg/interfaces"
	"go-vk-observer/internal/pkg/messages"
	"log"
	"strings"
)

const updateTimeout int = 10
const updateOffset int = 0

type Handler struct {
	telegramClient  interfaces.TelegramSenderInterface
	telegramService ServiceInterface
}

func NewHandler(
	telegramClient interfaces.TelegramSenderInterface,
	telegramService ServiceInterface,
) *Handler {
	return &Handler{
		telegramClient:  telegramClient,
		telegramService: telegramService,
	}
}

func (handler *Handler) HandleCommands() {
	var resultText string
	var err error
	u := tgbotapi.NewUpdate(updateOffset)
	u.Timeout = updateTimeout

	updates := handler.telegramClient.GetBot().GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		userName := update.Message.From.UserName
		telegramID := update.Message.Chat.ID
		arg := strings.TrimSpace(update.Message.CommandArguments())

		switch update.Message.Command() {
		case "start":
			resultText = handler.telegramService.Start(userName)
		case "add":
			resultText, err = handler.telegramService.AddSlug(telegramID, arg)
		case "delete":
			resultText, err = handler.telegramService.DeleteSlug(telegramID, arg)
		case "list":
			resultText, err = handler.telegramService.GetSlugsList(telegramID)
		default:
			resultText = messages.InvalidCommand
		}

		if err != nil {
			log.Println("Message error: ", err)
			resultText = messages.CommonError
		}

		err = handler.telegramClient.SendMessage(update.Message.Chat.ID, resultText, true)
		if err != nil {
			log.Println("Send message error: ", err)
			return
		}
	}
}
