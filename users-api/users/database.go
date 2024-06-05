package users

import (
	"errors"
	"log"
	"time"
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

func WriteNewUserRecord(
	new_user_id string,
	body *models.CreateUserRequest,
	params *Argon2idParams,
) (*models.User, error) {
	password_hash, err := GenerateHash(body.Password, params)
	if err != nil {
		return nil, err
	}

	new_record := models.User{
		ID:        new_user_id,
		Username:  body.Username,
		Email:     body.Email,
		Password:  password_hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := Database.Create(&new_record)
	if result.Error != nil {
		log.Println("Error occurred during user write: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("failed to save new user record to database")
	}

	return &new_record, nil
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

func WriteNewFollowingRecord(body *models.FollowUserRequest) (*models.Following, error) {
	new_record := models.Following{
		FollowerID: body.FollowerID,
		FollowedID: body.FollowedID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	result := Database.Create(&new_record)
	if result.Error != nil {
		log.Println("Error occurred during following write: " + result.Error.Error())
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, errors.New("failed to save new following record to database")
	}

	return &new_record, nil
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
