package services

import (
	"diluan/apis"
	"diluan/entities"
	"diluan/transform"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
)

var limit = 50
var requestLimit = "50"

type FacebookService interface {
	GetMessageDetail(messageID string, page entities.FacebookPage, fullUrl string) (apis.Message, error)
	GetConvoMessages(page entities.FacebookPage, convo entities.FacebookPageConversation) (apis.FBConvoMessages, error)
	GetMessageAttach(fullUrl string) (transform.Attachments, error)
}

type facebookService struct{}

func NewFacebookService() FacebookService {
	return &facebookService{}
}

func (f *facebookService) GetMessageDetail(messageID string, page entities.FacebookPage, fullUrl string) (apis.Message, error) {
	var fbMessage apis.Message
	if fullUrl == "" {
		fullUrl = os.Getenv("FACEBOOK_API_URL") + messageID + "?fields=id,created_time,from,to,message&access_token=" + page.AccessToken
	}
	fbMessagesRequest, _ := http.NewRequest("GET", fullUrl, nil)
	fbMessageResponse, fbMessageResponseError := http.DefaultClient.Do(fbMessagesRequest)
	if fbMessageResponseError != nil {
		return fbMessage, fbMessageResponseError
	}
	defer fbMessageResponse.Body.Close()
	responseData, err := ioutil.ReadAll(fbMessageResponse.Body)
	if err != nil {
		log.Fatal(err)
		return fbMessage, err
	}
	if err := json.Unmarshal(responseData, &fbMessage); err != nil {
		return fbMessage, err
	}
	return fbMessage, nil
}

func (f *facebookService) GetConvoMessages(page entities.FacebookPage, convo entities.FacebookPageConversation) (apis.FBConvoMessages, error) {
	var fbConvoMessages apis.FBConvoMessages
	count := convo.MessageCount
	ceil := float64(count) / float64(limit)
	ceilRequest := math.Ceil(ceil)

	numberRequest := 1
	if count > uint(limit) {
		numberRequest = int(ceilRequest)
	}
	fullUrl := os.Getenv("FACEBOOK_API_URL") + convo.ConversationID + "/messages?access_token=" + page.AccessToken + "&limit=" + requestLimit
	for i := 0; i < numberRequest; i++ {
		messagesLimit, err := HandleMessageWithLimit(page, convo.ConversationID, fullUrl)
		if err != nil {
			return apis.FBConvoMessages{}, err
		}
		if messagesLimit.Paging.Next != "" {
			fullUrl = messagesLimit.Paging.Next
		}
		fbConvoMessages.Data = append(fbConvoMessages.Data, messagesLimit.Data...)
	}
	return fbConvoMessages, nil

}

func HandleMessageWithLimit(page entities.FacebookPage, conversationID string, fullUrl string) (apis.FBConvoMessages, error) {
	var fbConvoMessages apis.FBConvoMessages
	fbConvoMessagesRequest, _ := http.NewRequest("GET", fullUrl, nil)
	fbConvoMessageResponse, fbConvoMessageResponseError := http.DefaultClient.Do(fbConvoMessagesRequest)
	if fbConvoMessageResponseError != nil {
		return apis.FBConvoMessages{}, fbConvoMessageResponseError
	}
	defer fbConvoMessageResponse.Body.Close()
	responseData, err := ioutil.ReadAll(fbConvoMessageResponse.Body)
	if err != nil {
		log.Fatal(err)
		return apis.FBConvoMessages{}, err
	}

	if err := json.Unmarshal(responseData, &fbConvoMessages); err != nil {
		return apis.FBConvoMessages{}, err
	}

	return fbConvoMessages, nil
}

func (f *facebookService) GetMessageAttach(fullUrl string) (transform.Attachments, error) {
	var messageAttachs transform.Attachments
	attachRequest, _ := http.NewRequest("GET", fullUrl, nil)
	attachResponse, attachResponseError := http.DefaultClient.Do(attachRequest)
	if attachResponseError != nil {
		return transform.Attachments{}, attachResponseError
	}
	defer attachResponse.Body.Close()

	responseData, err := ioutil.ReadAll(attachResponse.Body)

	if err != nil {
		log.Fatal(err)
		return transform.Attachments{}, err
	}

	if err := json.Unmarshal(responseData, &messageAttachs); err != nil {
		return transform.Attachments{}, err
	}

	return messageAttachs, nil
}
