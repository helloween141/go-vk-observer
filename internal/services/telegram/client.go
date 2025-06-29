package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-vk-observer/internal/pkg/messages"
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

func (c *Client) SendCommand(telegramID int64, text string) error {
	msg := tgbotapi.NewMessage(telegramID, text)
	msg.DisableWebPagePreview = true
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := c.bot.Send(msg)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *Client) SendVkPost(telegramID int64, slug string, title string, datetime string, text string) error {
	parseMode := tgbotapi.ModeMarkdown
	text = tgbotapi.EscapeText(parseMode, text)

	msg := tgbotapi.NewMessage(telegramID, fmt.Sprintf(messages.VkPostMessage, title, slug, datetime, text))
	msg.ParseMode = parseMode

	_, err := c.bot.Send(msg)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
