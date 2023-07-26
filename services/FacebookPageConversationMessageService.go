package services

import (
	"diluan/entities"
	"diluan/repositories"
)

type FacebookPageConversationMessageService interface {
	Create(messages []entities.FacebookPageConversationMessage, convertsationID string) []entities.FacebookPageConversationMessage
}

type facebookpageconversationmessageService struct {
	messageRepository repositories.FacebookPageConversationMessageRepository
}

func NewFacebookPageConversationMessageService(messageRepository repositories.FacebookPageConversationMessageRepository) FacebookPageConversationMessageService {
	return &facebookpageconversationmessageService{
		messageRepository: messageRepository,
	}
}

func (s *facebookpageconversationmessageService) Create(messages []entities.FacebookPageConversationMessage, convertsationID string) []entities.FacebookPageConversationMessage {
	interfaz_slice := ToInterfaceSlice(messages)
	s.messageRepository.Create(interfaz_slice, convertsationID)
	return []entities.FacebookPageConversationMessage{}
}

func ToInterfaceSlice(lists []entities.FacebookPageConversationMessage) []interface{} {
    iface := make([]interface{}, len(lists))
    for i := range lists {
        iface[i] = lists[i]
    }
    return iface
}
