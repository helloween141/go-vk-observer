package vk

import (
	"bytes"
	"embed"
	"fmt"
	"go-vk-observer/internal/pkg/utils"
	"go-vk-observer/internal/services/vk/responses"
	"html/template"
)

//go:embed templates/post_message.tmpl
var postTemplateFS embed.FS

type MessageData struct {
	GroupName  template.HTML
	PostUrl    template.HTML
	PostDate   string
	PostText   template.HTML
	RepostText template.HTML
}

type ServiceInterface interface {
	CreatePostMessage(groupName string, post responses.PostInfo) (string, error)
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) CreatePostMessage(groupName string, post responses.PostInfo) (string, error) {
	var repostText string

	postUrl := fmt.Sprintf("wall%d_%d", post.FromID, post.ID)
	postDate := utils.FormatTimestampToDatetime(post.Date)
	postText := utils.FormatPostLinks(post.Text)

	if post.RepostInfo != nil {
		repostText = utils.FormatRepostLinks(post.RepostInfo[0].Text)
	}

	text, err := renderPostText(
		MessageData{
			GroupName:  template.HTML(groupName),
			PostUrl:    template.HTML(postUrl),
			PostDate:   postDate,
			PostText:   template.HTML(postText),
			RepostText: template.HTML(repostText),
		})

	if err != nil {
		return "", err
	}

	return text, nil
}

func renderPostText(data MessageData) (string, error) {
	tmpl, err := template.ParseFS(postTemplateFS, "templates/post_message.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
