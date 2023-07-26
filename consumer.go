package main

import (
	"diluan/config"
	"diluan/database"
	"diluan/entities"
	"diluan/repositories"
	"diluan/services"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB        = database.InitMysql()
	mongodb *mongo.Database = config.InitMongoDB()

	//Repository
	userRepository               repositories.UserRepository                            = repositories.NewUserRepository(db)
	tokenRepository              repositories.OauthAccessTokenRepository                = repositories.NewOauthAccessTokenRepository(db)
	facebookPageRepository       repositories.FacebookPageRepository                    = repositories.NewFacebookPageRepository(db)
	fbPageConvoRepository        repositories.FacebookPageConversationRepository        = repositories.NewFacebookPageConversationRepository(db)
	messageRepository            repositories.FacebookPageConversationMessageRepository = repositories.NewFacebookPageConversationMessageRepository(db, mongodb)
	fbPageConvoMessageRepository repositories.FacebookPageConversationMessageRepository = repositories.NewFacebookPageConversationMessageRepository(db , mongodb)

	//Service
	jwtService          services.JWTService                             = services.NewJWTService()
	facebookService     services.FacebookService                        = services.NewFacebookService()
	facebookPageService services.FacebookPageService                    = services.NewFacebookPageService(facebookPageRepository, fbPageConvoRepository, fbPageConvoMessageRepository)
	authService         services.AuthService                            = services.NewAuthService(userRepository, tokenRepository)
	userService         services.UserService                            = services.NewUserService(userRepository)
	queueService        services.QueueService                           = services.NewQueueService()
	messageService      services.FacebookPageConversationMessageService = services.NewFacebookPageConversationMessageService(messageRepository)

	queueName     = "LUMIA_QUEUE"
	layout        = "2006-01-02T15:04:05+0000"
	attachmentUrl = ""
)

func main() {
	fmt.Println("Consumer App")
	godotenv.Load()
	amqpServer := os.Getenv("AMQP_URL")
	conn, err := amqp.Dial(amqpServer)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"LUMIA_QUEUE",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for p := range msgs {
			body := p.Body
			headers := p.Headers
			fmt.Println(headers)
			//############ Start Handle Sync Facebook Conversation
			go HandleFacebookPageConversation(headers, body)
			go HandleConversationDetail(headers, body)
			//############ End Handle Sync Facebook Conversation
		}
	}()

	fmt.Println("Success")
	fmt.Println("[*] - waiting for message")
	<-forever
}

func HandleFacebookPageConversation(headers map[string]interface{}, body []byte) {
	if headers["Sync_Conversation_From_Page"] == nil {
		return
	}
	var convos []entities.FacebookPageConversation
	if err := json.Unmarshal(body, &convos); err != nil {
		panic(err)
	}

	for _, c := range convos {
		fmt.Printf("MESSAGE COUNT %v", c.MessageCount)
		convo, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		mapHeaders := make(map[string]interface{})
		mapHeaders["Sync_Conversation_Detail"] = "convo"
		go queueService.MakePublish(queueName, convo, mapHeaders)
	}

}

func HandleConversationDetail(headers map[string]interface{}, body []byte) {

	if headers["Sync_Conversation_Detail"] == nil {
		return
	}
	var convo entities.FacebookPageConversation
	if err := json.Unmarshal(body, &convo); err != nil {
		panic(err)
	}

	page := facebookPageRepository.FindByPageID(convo.PageID)
	fbConvoMessages, err := facebookService.GetConvoMessages(page, convo)
	if err != nil {
		panic(err)
	}
	var messagesEntity []entities.FacebookPageConversationMessage
	for _, message := range fbConvoMessages.Data {
		messageDetail, err := facebookService.GetMessageDetail(message.ID, page, "")
		if err != nil {
			panic(err)
		}
		updatedTime, _ := time.Parse(layout, messageDetail.CreatedTime)
		singleMessage := entities.FacebookPageConversationMessage{
			PageID:         page.PageID,
			ConversationID: convo.ConversationID,
			MessageID:      messageDetail.ID,
			Message:        messageDetail.Message,
			From:           messageDetail.From.Name,
			To:             messageDetail.To.Data[0].Name,
			Sender:         messageDetail.From.ID,
			CreatedTime:    updatedTime,
		}

		if messageDetail.Message == "" {
			fullUrl := os.Getenv("FACEBOOK_API_URL") + singleMessage.MessageID + "/attachments?access_token=" + page.AccessToken
			attachment, _ := facebookService.GetMessageAttach(fullUrl)
			if len(attachment.Data) > 0 {
				singleMessage.AttachmentID = attachment.Data[0].ID
				singleMessage.AttachmentName = attachment.Data[0].Name
				singleMessage.AttachmentSize = attachment.Data[0].Size
				singleMessage.AttachmentType = attachment.Data[0].MimeType
				singleMessage.AttachmentUrl = GetAttachmentUrl(attachment.Data[0].MimeType, attachment.Data[0].FileUrl, attachment.Data[0].VideoData.Url)
			}
		}
		messagesEntity = append(messagesEntity, singleMessage)
		// fmt.Println(singleMessage.ID)
	}
	messageService.Create(messagesEntity, convo.ConversationID)

	return

}

func GetAttachmentUrl(mimeType string, fileUrl string, url string) string {
	if mimeType == "audio/mp4" {
		return fileUrl
	}
	return url
}
