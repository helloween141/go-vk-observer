package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-vk-observer/internal/pkg/messages"
)

type Client struct {
	bot *tgbotapi.BotAPI
}

func NewClient(token string) (*Client, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Client{bot: bot}, err
}

func (c *Client) GetBot() *tgbotapi.BotAPI {
	return c.bot
}

func (c *Client) SendMessage(telegramID int64, text string, isCommand bool) error {
	msg := tgbotapi.NewMessage(telegramID, text)

	if isCommand {
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.DisableWebPagePreview = true
	} else {
		msg.ParseMode = tgbotapi.ModeHTML
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendMenu(telegramID int64) error {
	msg := tgbotapi.NewMessage(telegramID, messages.MenuMessage)
	msg.ParseMode = tgbotapi.ModeMarkdown

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Добавить", "add"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Список", "list"),
		),
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
