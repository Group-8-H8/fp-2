package database

import (
	"fmt"
	"github.com/Group-8-H8/fp-2.git/config"
	"github.com/Group-8-H8/fp-2.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() error {
	conn := fmt.Sprintf("host=%s  user=%s password=%s dbname=%s port=%s sslmode=disable", config.HOST, config.USERNAME, config.PASSWORD, config.DB_NAME, config.PORT)
	db, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Successfully Connected to Database: ", config.DB_NAME)
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}