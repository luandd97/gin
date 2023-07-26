package repositories

import (
	"diluan/entities"

	"gorm.io/gorm"
)

type FacebookPageConversationRepository interface {
	Create(convoEntities []entities.FacebookPageConversation) []entities.FacebookPageConversation
	FindByPageIDAndConvoID(pageID string, convoID string) entities.FacebookPageConversation
	GetByPageIDAndUserID(pageID string, userID uint64) []entities.FacebookPageConversation
}

type facebookPageConversationConnection struct {
	connection *gorm.DB
}

func NewFacebookPageConversationRepository(db *gorm.DB) FacebookPageConversationRepository {
	return &facebookPageConversationConnection{
		connection: db,
	}
}

func (db *facebookPageConversationConnection) Create(convoEntities []entities.FacebookPageConversation) []entities.FacebookPageConversation {
	db.connection.Create(&convoEntities)
	return convoEntities
}

func (db *facebookPageConversationConnection) FindByPageIDAndConvoID(pageID string, convoID string) entities.FacebookPageConversation {
	var convo entities.FacebookPageConversation
	db.connection.Where("page_id = ?", pageID).
		Where("conversation_id = ?", convoID).
		Take(&convo)
	return convo
}

func (db *facebookPageConversationConnection) GetByPageIDAndUserID(pageID string, userID uint64) []entities.FacebookPageConversation {
	var convos []entities.FacebookPageConversation
	db.connection.Where("page_id = ?", pageID).
		Where("user_id = ?", userID).
		Find(&convos)
	return convos
}
