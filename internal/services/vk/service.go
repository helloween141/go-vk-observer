package vk

import (
	"fmt"
	"go-vk-observer/internal/pkg/messages"
	"go-vk-observer/internal/pkg/utils"
	"go-vk-observer/internal/services/vk/responses"
)

type ServiceInterface interface {
	GetPostText(groupName string, post responses.PostInfo) string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetPostText(groupName string, post responses.PostInfo) string {
	var repostText string

	postUrl := fmt.Sprintf("wall%d_%d", post.FromID, post.ID)
	postDate := utils.FormatTimestampToDatetime(post.Date)
	postText := utils.FormatPostLinks(post.Text)

	if post.RepostInfo != nil {
		repostText = utils.FormatRepostLinks(post.RepostInfo[0].Text)
	}

	return fmt.Sprintf(
		messages.VkPostMessage,
		groupName,
		postUrl,
		postDate,
		postText,
		repostText,
	)
}
