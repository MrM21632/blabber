package posts

import (
	"errors"
	"posts/posts/models"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Post{})
	Database = database
}

func WriteNewPostRecord(
	new_post_id string,
	body *models.CreatePostRequest,
) (*models.Post, error) {
	new_record := models.Post{
		ID:        new_post_id,
		ParentID:  body.ParentID,
		UserID:    body.UserID,
		Contents:  body.Contents,
		Likes:     0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Infof("%+v", new_record)
	result := Database.Create(&new_record)
	if result.Error != nil {
		log.Error("Error occurred during user write: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("failed to save new post record to database")
	}

	return &new_record, nil
}
