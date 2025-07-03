package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
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
	msg.ParseMode = tgbotapi.ModeMarkdown

	if isCommand {
		msg.DisableWebPagePreview = true
	}

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
