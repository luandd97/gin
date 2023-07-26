package repositories

import (
	"diluan/entities"

	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(tokenEntities entities.OauthAccessToken) entities.OauthAccessToken
	Update(tokenEntities entities.OauthAccessToken) entities.OauthAccessToken
	FindByUserID(userID uint64) entities.OauthAccessToken
}

type oauthaccesstokenConnection struct {
	connection *gorm.DB
}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &oauthaccesstokenConnection{
		connection: db,
	}
}

func (db *oauthaccesstokenConnection) Create(tokenEntities entities.OauthAccessToken) entities.OauthAccessToken {
	db.connection.Save(&tokenEntities)
	return tokenEntities
}

func (db *oauthaccesstokenConnection) Update(tokenEntities entities.OauthAccessToken) entities.OauthAccessToken {
	db.connection.Updates(&tokenEntities)
	db.connection.Find(&tokenEntities)
	return tokenEntities
}

func (db *oauthaccesstokenConnection) FindByUserID(userID uint64) entities.OauthAccessToken {
	var token entities.OauthAccessToken
	db.connection.
		Where("user_id = ?", userID).
		Take(&token)
	return token
}
