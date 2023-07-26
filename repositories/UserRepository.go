package repositories

import (
	"diluan/entities"
	"log"
	"strings"

	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entities.User) entities.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.User
	FindById(id uint64) entities.User
	FindByEmailAndDriver(email string, driver string) entities.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmailAndDriver(email string, driver string) entities.User {
	var user entities.User
	db.connection.
		Where("email = ?", email).
		Where("social_driver = ?", driver).
		Take(&user)
	return user
}

func (db *userConnection) Create(user entities.User) entities.User {
	user.Password = hashAndSalt([]byte(user.Password))
	token := ksuid.New()
	user.Token = strings.ToUpper(string(token.String()))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Find(&user)
	return user
}

func (db *userConnection) FindById(id uint64) entities.User {
	var user entities.User
	db.connection.Where("id = ?", id).Find(&user)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
