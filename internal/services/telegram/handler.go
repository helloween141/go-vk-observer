package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-vk-observer/internal/pkg/interfaces"
	"go-vk-observer/internal/pkg/messages"
	"go-vk-observer/internal/pkg/utils"
	"log"
)

const updateTimeout int = 10
const updateOffset int = 0

const (
	StateDefault = iota
	StateWaitingForAdd
	StateWaitingForDelete
	StateList
)

type UserState struct {
	State int
}

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
	var telegramID int64
	var err error

	userStates := make(map[int64]UserState)

	u := tgbotapi.NewUpdate(updateOffset)
	u.Timeout = updateTimeout

	updates := handler.telegramClient.GetBot().GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			telegramID = update.CallbackQuery.Message.Chat.ID

			handleCallbackQuery(update.CallbackQuery, handler, userStates)

			if userStates[telegramID].State == StateList {
				resultText, _ = handler.telegramService.GetSlugsList(telegramID)
			}
		}

		if update.Message != nil {
			telegramID = update.Message.Chat.ID
			userName := update.Message.From.UserName
			slug := utils.ExtractVkSlug(update.Message.Text)

			switch userStates[telegramID].State {
			case StateWaitingForAdd:
				resultText, _ = handler.telegramService.AddSlug(telegramID, slug)
			case StateWaitingForDelete:
				resultText, _ = handler.telegramService.DeleteSlug(telegramID, slug)
			default:
				resultText = handler.telegramService.Start(userName)
			}
		}

		if resultText != "" {
			err = handler.telegramClient.SendMessage(telegramID, resultText, true)
			if err != nil {
				log.Println("Send message error: ", err)
				return
			}

			_ = handler.telegramClient.SendMenu(telegramID)

			resultText = ""
			userStates[telegramID] = UserState{State: StateDefault}
		}
	}
}

func handleCallbackQuery(query *tgbotapi.CallbackQuery, handler *Handler, userStates map[int64]UserState) {
	var resultText string
	var err error
	telegramID := query.Message.Chat.ID

	switch query.Data {
	case "add":
		resultText = messages.AddSlugTooltip
		userStates[telegramID] = UserState{State: StateWaitingForAdd}
	case "delete":
		resultText = messages.DeleteSlugTooltip
		userStates[telegramID] = UserState{State: StateWaitingForDelete}
	case "list":
		userStates[telegramID] = UserState{State: StateList}
	}

	if resultText != "" {
		err = handler.telegramClient.SendMessage(telegramID, resultText, true)
		if err != nil {
			log.Println("Send message error: ", err)
			return
		}
	}
}
