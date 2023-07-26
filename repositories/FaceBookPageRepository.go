package repositories

import (
	"diluan/entities"

	"gorm.io/gorm"
)

type FacebookPageRepository interface {
	Create(fbPageEntity entities.FacebookPage) entities.FacebookPage
	CreateBatch(fbPageEntity []entities.FacebookPage) []entities.FacebookPage
	GetByUserId(userID uint64) []entities.FacebookPage
	FindByPageID(pageID string) entities.FacebookPage
	DeleteByUserID(userID uint64)
	Update(page entities.FacebookPage) entities.FacebookPage
}

type facebookpageConnection struct {
	connection *gorm.DB
}

func NewFacebookPageRepository(db *gorm.DB) FacebookPageRepository {
	return &facebookpageConnection{
		connection: db,
	}
}

func (db *facebookpageConnection) GetByUserId(userID uint64) []entities.FacebookPage {
	var fbPages []entities.FacebookPage
	db.connection.Where("user_id = ?", userID).Find(&fbPages)
	return fbPages
}

func (db *facebookpageConnection) FindByPageID(pageID string) entities.FacebookPage {
	var fbPage entities.FacebookPage
	db.connection.Where("page_id = ?", pageID).Take(&fbPage)
	return fbPage
}

func (db *facebookpageConnection) Create(fbPageEntity entities.FacebookPage) entities.FacebookPage {
	db.connection.Save(&fbPageEntity)
	return fbPageEntity
}

func (db *facebookpageConnection) CreateBatch(fbPageEntity []entities.FacebookPage) []entities.FacebookPage {
	db.connection.Create(&fbPageEntity)
	return fbPageEntity
}

func (db *facebookpageConnection) DeleteByUserID(userID uint64) {
	var fbPageEntity []entities.FacebookPage
	db.connection.Where("user_id = ?", userID).Find(&fbPageEntity)
	db.connection.Where("user_id = ?", userID).Delete(&fbPageEntity)
}

func (db *facebookpageConnection) Update(page entities.FacebookPage) entities.FacebookPage {
	db.connection.Updates(&page)
	db.connection.Find(&page)
	return page
}
