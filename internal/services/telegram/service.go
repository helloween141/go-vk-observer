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
	GetSlugsList(telegramID int64) (string, error)
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

	vkEntity, _ = service.vkRepository.GetBySlug(slug)

	if vkEntity == nil {
		var name string

		groupResponse, errGroupResponse := service.vkClient.SendGetGroupRequest(slug)
		if errGroupResponse != nil {
			return "", errGroupResponse
		}

		groupInfo := groupResponse.Response.Groups
		if groupInfo == nil {
			userResponse, errUserResponse := service.vkClient.SendGetUserRequest(slug)
			if errUserResponse != nil {
				return fmt.Sprintf(messages.SlugNotFound, slug), nil
			}

			userInfo := userResponse.Response
			name = fmt.Sprintf("%s %s", userInfo[0].FirstName, userInfo[0].LastName)
		} else {
			name = groupInfo[0].Name
		}

		vkEntity, err = service.vkRepository.Create(slug, name)
		if err != nil {
			return "", err
		}
	}

	if service.telegramRepository.IsEntityExistsByTelegramID(telegramID, vkEntity.ID) {
		return fmt.Sprintf(messages.SlugAlreadyExists, slug), nil
	}

	err = service.telegramRepository.Create(telegramID, vkEntity.ID)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(messages.AddSlugSuccessful, slug), nil
}

func (service *Service) DeleteSlug(telegramID int64, slug string) (string, error) {
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

func (service *Service) GetSlugsList(telegramID int64) (string, error) {
	var resultText string

	entities, err := service.telegramRepository.GetByTelegramID(telegramID)
	if err != nil {
		return "", err
	}

	if len(entities) == 0 {
		return messages.SlugsListIsEmpty, nil
	}

	resultText = messages.SlugsListTitle
	for i, entity := range entities {
		resultText += fmt.Sprintf(messages.SlugsListMessage, i+1, entity.Name, entity.Slug, entity.Slug)
	}

	return resultText, nil
}
