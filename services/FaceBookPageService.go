package services

import (
	"diluan/apis"
	"diluan/dto"
	"diluan/entities"
	"diluan/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mashingan/smapping"
)

type FacebookPageService interface {
	HandleFacebookPage(token string, fbUserId string, userId uint64)
	GetByUserId(userID uint64) []entities.FacebookPage
	FindByPageID(pageID string) entities.FacebookPage
	GetPageConversation(token string, pageID string, userID uint64) []entities.FacebookPageConversation
	Update(page entities.FacebookPage) entities.FacebookPage
	FindByPageIDAndConvoID(pageID string, convoID string) entities.FacebookPageConversation
	GetByPageIDAndUserID(pageID string, userID uint64) []entities.FacebookPageConversation
	GetPageConversationDetail(page entities.FacebookPage, convo entities.FacebookPageConversation) []entities.FacebookPageConversationMessage
}

type facebookpageService struct {
	facebookPageRepository             repositories.FacebookPageRepository
	facebookPageConvoRepository        repositories.FacebookPageConversationRepository
	facebookPageConvoMessageRepository repositories.FacebookPageConversationMessageRepository
}

func NewFacebookPageService(facebookPageRepository repositories.FacebookPageRepository, facebookPageConvoRepository repositories.FacebookPageConversationRepository, facebookPageConvoMessageRepository repositories.FacebookPageConversationMessageRepository) FacebookPageService {
	return &facebookpageService{
		facebookPageRepository:             facebookPageRepository,
		facebookPageConvoRepository:        facebookPageConvoRepository,
		facebookPageConvoMessageRepository: facebookPageConvoMessageRepository,
	}
}

func (s *facebookpageService) GetByUserId(userID uint64) []entities.FacebookPage {
	return s.facebookPageRepository.GetByUserId(userID)
}

func (s *facebookpageService) HandleFacebookPage(token string, fbUserId string, userId uint64) {
	var facebookPage apis.FacebookPageResponse
	s.facebookPageRepository.DeleteByUserID(userId)
	fullUrl := os.Getenv("FACEBOOK_API_URL") + fbUserId + "/accounts?fields=picture,access_token,name&access_token=" + token
	fbPageRequest, _ := http.NewRequest("GET", fullUrl, nil)
	fbPageResponse, fbPageResponseError := http.DefaultClient.Do(fbPageRequest)
	if fbPageResponseError != nil {
		panic("adasdsadsadsad")
	}
	defer fbPageResponse.Body.Close()
	responseData, err := ioutil.ReadAll(fbPageResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(responseData, &facebookPage)
	sc := []entities.FacebookPage{}
	for _, page := range facebookPage.Data {

		fbPageDto := dto.FacebookPageDTO{
			UserID:      userId,
			PageID:      page.ID,
			AccessToken: page.AccessToken,
			Name:        page.Name,
			Picture:     page.Picture.Data.Url,
		}
		fbPageEntity := entities.FacebookPage{}
		err := smapping.FillStruct(&fbPageEntity, smapping.MapFields(fbPageDto))

		if err != nil {
			log.Fatal("Failed Map %v", err)
		}
		sc = append(sc, fbPageEntity)

	}
	s.facebookPageRepository.CreateBatch(sc)

}

func (s *facebookpageService) FindByPageID(pageID string) entities.FacebookPage {
	return s.facebookPageRepository.FindByPageID(pageID)
}

func (s *facebookpageService) GetPageConversation(token string, pageID string, userID uint64) []entities.FacebookPageConversation {
	var conversations apis.FacebookPageConversation
	fullUrl := os.Getenv("FACEBOOK_API_URL") + pageID + "/conversations?fields=snippet,subject,senders,message_count,unread_count,link,updated_time&limit=100000&access_token=" + token
	fbConvoRequest, _ := http.NewRequest("GET", fullUrl, nil)
	convoResponse, convoResponseError := http.DefaultClient.Do(fbConvoRequest)
	if convoResponseError != nil {
		panic("adasdsadsadsad")
	}

	responseData, err := ioutil.ReadAll(convoResponse.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, &conversations)

	convoEntities := []entities.FacebookPageConversation{}
	layout := "2006-01-02T15:04:05+0000"

	for _, convo := range conversations.Data {
		fmt.Println(convo.Snippet)
		str := convo.UpdatedTime
		updatedTime, _ := time.Parse(layout, str)
		convoDTO := dto.FacebookPageConversationDTO{
			UserID:         userID,
			PageID:         pageID,
			ConversationID: convo.ID,
			InboxUrl:       convo.Link,
			IsSync:         false,
			LastMessage:    convo.Snippet,
			UnreadCount:    convo.UnreadCount,
			UpdatedTime:    updatedTime,
			Sender:         convo.Senders.Data[0].Name,
			MessageCount:   uint(convo.MessageCount),
		}

		convoEntity := entities.FacebookPageConversation{}
		err := smapping.FillStruct(&convoEntity, smapping.MapFields(convoDTO))
		if err != nil {
			log.Fatal("Failed Map %v", err)
		}
		convoEntities = append(convoEntities, convoEntity)
	}
	convos := s.facebookPageConvoRepository.Create(convoEntities)
	defer convoResponse.Body.Close()
	return convos
}

func (s *facebookpageService) Update(page entities.FacebookPage) entities.FacebookPage {
	return s.facebookPageRepository.Update(page)
}

func (s *facebookpageService) FindByPageIDAndConvoID(pageID string, convoID string) entities.FacebookPageConversation {
	return s.facebookPageConvoRepository.FindByPageIDAndConvoID(pageID, convoID)
}

func (s *facebookpageService) GetPageConversationDetail(page entities.FacebookPage, convo entities.FacebookPageConversation) []entities.FacebookPageConversationMessage {
	return s.facebookPageConvoMessageRepository.GetPageConversationDetail(convo.PageID, convo.ConversationID)
}

func (s *facebookpageService) GetByPageIDAndUserID(pageID string, userID uint64) []entities.FacebookPageConversation {
	return s.facebookPageConvoRepository.GetByPageIDAndUserID(pageID, userID)
}
