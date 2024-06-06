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
		Handle:    body.Handle,
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

func UpdateUserRecord(body *models.UpdateUserRequest) error {
	to_update := models.User{}
	query := Database.First(&to_update, "id = ?", body.ID)
	if query.Error != nil {
		log.Println("Error occurred during find: " + query.Error.Error())
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("record not found in database")
	}

	if body.Handle != nil {
		to_update.Handle = *body.Handle
	}
	if body.Username != nil {
		to_update.Username = *body.Username
	}
	if body.Email != nil {
		to_update.Email = *body.Email
	}
	to_update.UpdatedAt = time.Now()

	result := Database.Save(&to_update)
	if result.Error != nil {
		log.Println("Error occurred during update: " + result.Error.Error())
		return result.Error
	}
	return nil
}

func UpdatePasswordHashForUser(user_id, new_password string, params *Argon2idParams) error {
	to_update := models.User{}
	query := Database.First(&to_update, "id = ?", user_id)
	if query.Error != nil {
		log.Println("Error occurred during find: " + query.Error.Error())
		return query.Error
	}
	if query.RowsAffected == 0 {
		return errors.New("record not found in database")
	}

	password_hash, err := GenerateHash(new_password, params)
	if err != nil {
		return err
	}
	to_update.Password = password_hash

	result := Database.Save(&to_update)
	if result.Error != nil {
		log.Println("Error occurred during password update: " + result.Error.Error())
		return result.Error
	}
	return nil
}

func DeleteUserRecord(user_id string) error {
	var record models.User
	found := Database.Where("user_id = ?", user_id).First(&record)
	if found.Error != nil {
		log.Println("Error occurred during find: " + found.Error.Error())
		return found.Error
	}
	if found.RowsAffected == 0 {
		return errors.New("record not found in database")
	}

	result := Database.Delete(&record)
	if result.Error != nil {
		log.Println("Error occurred during delete: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record already deleted")
	}
	return nil
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

func GetFollowersForUser(followed_id string) ([]models.UserFollowerEntity, error) {
	var followers []models.UserFollowerEntity
	query := Database.
		Table("user").
		Select("user.id, user.handle, user.username").
		Joins("inner join user_following on users.id = user_following.follower_id").
		Where("user_following.followed_id = ?", followed_id).
		Scan(&followers)
	if query.Error != nil {
		log.Println("Error occurred during query: " + query.Error.Error())
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return []models.UserFollowerEntity{}, nil
	}

	return followers, nil
}

func GetFollowsForUser(follower_id string) ([]models.UserFollowerEntity, error) {
	var follows []models.UserFollowerEntity
	query := Database.
		Table("user").
		Select("user.id, user.handle, user.username").
		Joins("inner join user_following on users.id = user_following.followed_id").
		Where("user_following.follower_id = ?", follower_id).
		Scan(&follows)
	if query.Error != nil {
		log.Println("Error occurred during query: " + query.Error.Error())
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return []models.UserFollowerEntity{}, nil
	}

	return follows, nil
}

func DeleteFollowingRecord(follower_id, followed_id string) error {
	var record models.Following
	found := Database.
		Where("user_following.follower_id = ?", follower_id).
		Where("user_following.followed_id = ?", followed_id).
		First(&record)
	if found.Error != nil {
		log.Println("Error occurred during find: " + found.Error.Error())
		return found.Error
	}
	if found.RowsAffected == 0 {
		return errors.New("record not found in database")
	}

	result := Database.Delete(&record)
	if result.Error != nil {
		log.Println("Error occurred during delete: " + result.Error.Error())
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("record already deleted")
	}
	return nil
}
