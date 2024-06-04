package users

import (
	"errors"
	"log"
	"users/users/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(
		&models.User{},
		&models.Following{},
	)
	Database = database
}

func GetUserRecord(user_id string) (*models.User, error) {
	var record models.User
	result := Database.Where("user.id = ?", user_id).First(&record)
	if result.Error != nil {
		log.Println("Error occurred during query: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found in database")
	}

	return &record, nil
}

func GetFollowingRecord(follower_id, followed_id string) (*models.Following, error) {
	var record models.Following
	result := Database.
		Where("user_following.follower_id = ?", follower_id).
		Where("user_following.followed_id = ?", followed_id).
		First(&record)
	if result.Error != nil {
		log.Println("Error occurred during query: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("record not found in database")
	}

	return &record, nil
}
