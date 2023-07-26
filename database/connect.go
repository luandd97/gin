package database

import (
	"diluan/entities"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func InitMysql() *gorm.DB {
	db = connectDB()
	return db
}

func connectDB() *gorm.DB {
	errConnect := godotenv.Load()
	if errConnect != nil {
		log.Fatalf("err loading: %v", errConnect)
	}

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	db.AutoMigrate(
		&entities.Role{},
		&entities.User{},
		&entities.OauthAccessToken{},
		&entities.FacebookPage{},
		&entities.FacebookPageConversation{},
		&entities.FacebookPageConversationMessage{},
	)
	// if condition {

	// }
	Migrate(db)
	return db
}
