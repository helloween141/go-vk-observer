package telegram

import (
	"fmt"
	"go-vk-observer/internal/database/gen/dbstore"
	"go-vk-observer/internal/pkg/interfaces"
	"go-vk-observer/internal/pkg/messages"
)

type ServiceInterface interface {
	Start(userName string) string
	AddSlug(telegramID int64, slug string) (string, error)
	DeleteSlug(telegramID int64, slug string) (string, error)
	GetAllSlugs(telegramID int64) (string, error)
}

type Service struct {
	telegramSender     interfaces.TelegramSenderInterface
	vkClient           interfaces.VkClientInterface
	telegramRepository interfaces.TelegramRepositoryInterface
	vkRepository       interfaces.VkRepositoryInterface
}

func NewService(
	telegramSender interfaces.TelegramSenderInterface,
	vkClient interfaces.VkClientInterface,
	telegramRepository interfaces.TelegramRepositoryInterface,
	vkRepository interfaces.VkRepositoryInterface,
) *Service {
	return &Service{
		telegramRepository: telegramRepository,
		vkRepository:       vkRepository,
		telegramSender:     telegramSender,
		vkClient:           vkClient,
	}
}

func (service *Service) Start(userName string) string {
	return fmt.Sprintf(messages.StartMessage, userName)
}

func (service *Service) AddSlug(telegramID int64, slug string) (string, error) {
	var vkEntity *dbstore.VkEntity
	var err error

	if slug == "" {
		return messages.SlugIsEmpty, nil
	}

	vkEntity, err = service.vkRepository.GetBySlug(slug)
	if vkEntity == nil {
		response, errReq := service.vkClient.SendGetGroupRequest(slug)
		if errReq != nil {
			return "", errReq
		}

		groupInfo := response.Response.Groups
		if groupInfo == nil {
			return "", err
		}

		vkEntity, err = service.vkRepository.Create(slug, groupInfo[0].Name, string(dbstore.EntityTypeWALL))
		if err != nil {
			return "", err
		}
	} else {
		return fmt.Sprintf(messages.SlugAlreadyExists, slug), nil
	}

	err = service.telegramRepository.Create(telegramID, vkEntity.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(messages.AddSlugSuccessful, slug), nil
}

func (service *Service) DeleteSlug(telegramID int64, slug string) (string, error) {
	if slug == "" {
		return messages.SlugIsEmpty, nil
	}

	vkEntity, _ := service.vkRepository.GetBySlug(slug)
	if vkEntity == nil {
		return fmt.Sprintf(messages.SlugNotFound, slug), nil
	}

	err := service.telegramRepository.Delete(telegramID, vkEntity.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(messages.DeleteSlugSuccessful, slug), nil
}

func (service *Service) GetAllSlugs(telegramID int64) (string, error) {
	var resultText string

	entities, err := service.telegramRepository.GetByTelegramID(telegramID)
	if err != nil {
		return "", err
	}

	for i, entity := range entities {
		resultText += fmt.Sprintf(messages.SlugsListMessage, i+1, entity.Name, entity.Slug, entity.Slug)
	}

	return resultText, nil
}
