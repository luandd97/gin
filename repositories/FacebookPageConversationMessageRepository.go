package repositories

import (
	"context"
	"diluan/entities"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
)

type FacebookPageConversationMessageRepository interface {
	Create(messages []interface{}, convertsationID string)
	GetPageConversationDetail(pageID string, conversationID string) []entities.FacebookPageConversationMessage
}

type facebookpageconversationmessageConnection struct {
	connection *gorm.DB
	mongodb    *mongo.Database
}

func NewFacebookPageConversationMessageRepository(db *gorm.DB, mongodb *mongo.Database) FacebookPageConversationMessageRepository {
	return &facebookpageconversationmessageConnection{
		connection: db,
		mongodb:    mongodb,
	}
}
func (db *facebookpageconversationmessageConnection) Create(messages []interface{}, convertsationID string) {
	collection := db.mongodb.Collection(convertsationID)

	res, _ := collection.InsertMany(context.TODO(), messages)
	fmt.Println(res)
}

func (db *facebookpageconversationmessageConnection) GetPageConversationDetail(pageID string, conversationID string) []entities.FacebookPageConversationMessage {
	var messages []entities.FacebookPageConversationMessage
	// db.connection.Where("page_id = ?", pageID).
	// 	Where("conversation_id = ?", conversationID).
	// 	Find(&messages)
	opts := options.Find().SetSort(bson.D{{Key: "created_time", Value: 1}})
	cur, err := db.mongodb.Collection(conversationID).Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		bsonByte, _ := bson.Marshal(result)
		var message entities.FacebookPageConversationMessage
		_ = bson.Unmarshal(bsonByte, &message)
		messages = append(messages, message)
		// do something with result....
	}
	fmt.Println(err)
	return messages
}
