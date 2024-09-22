package api

import (
	"time"
	"users-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

func WriteNewUserRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.CreateUserRequest,
	password_hash string,
	user_id uuid.UUID,
) (pgconn.CommandTag, error) {
	var bio_text string
	if input.Bio == nil {
		bio_text = ""
	} else {
		bio_text = *input.Bio
	}

	query_string := `
	INSERT INTO blabber.user (
		id, username, user_handle, user_bio, email, password_hash, created_at, updated_at,
		followers, follows
	) VALUES (
		@id, @username, @handle, @bio, @email, @hash, @createdAt, @updatedAt, @followers,
		@follows
	);
	`
	query_args := pgx.NamedArgs{
		"id":        user_id,
		"username":  input.Username,
		"handle":    input.Handle,
		"bio":       bio_text,
		"email":     input.Email,
		"hash":      password_hash,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
		"followers": 0,
		"follows":   0,
	}

	return pool.Exec(context, query_string, query_args)
}

func RetrieveUserRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) (*models.User, error) {
	query_string := `
	SELECT * FROM blabber.user
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	row, err := pool.Query(context, query_string, query_args)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.User])
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUserRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.UpdateUserRequest,
	user models.User,
) (pgconn.CommandTag, error) {
	var username string
	if input.Bio == nil {
		username = user.Username
	} else {
		username = *input.Username
	}

	var handle string
	if input.Bio == nil {
		handle = user.Handle
	} else {
		handle = *input.Handle
	}

	var email string
	if input.Bio == nil {
		email = user.Email
	} else {
		email = *input.Email
	}

	var bio_text string
	if input.Bio == nil {
		bio_text = user.Bio
	} else {
		bio_text = *input.Bio
	}

	query_string := `
	UPDATE blabber.user
	SET
		username = @username,
		user_handle = @handle,
		email = @email,
		user_bio = @bio,
		updated_at = @updated_at
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id":         input.ID,
		"username":   username,
		"handle":     handle,
		"email":      email,
		"bio":        bio_text,
		"updated_at": time.Now(),
	}

	return pool.Exec(context, query_string, query_args)
}

func DeleteUserRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) (pgconn.CommandTag, error) {
	query_string := `
	DELETE FROM blabber.user
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	return pool.Exec(context, query_string, query_args)
}

func RetrieveUserPassword(
	context *gin.Context,
	pool *pgxpool.Pool,
	user_id uuid.UUID,
) (*string, error) {
	query_string := `
	SELECT password_hash FROM blabber.user
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id": user_id,
	}

	var password_hash string
	err := pool.QueryRow(context, query_string, query_args).Scan(&password_hash)
	if err != nil {
		return nil, err
	}

	return &password_hash, nil
}

func UpdateUserPassword(
	context *gin.Context,
	pool *pgxpool.Pool,
	user_id uuid.UUID,
	new_password_hash string,
) (pgconn.CommandTag, error) {
	query_string := `
	UPDATE blabber.user
	SET
		password_hash = @password,
		updated_at = @updated_at
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id":         user_id,
		"password":   new_password_hash,
		"updated_at": time.Now(),
	}

	return pool.Exec(context, query_string, query_args)
}

func RetrieveFollowerRecordsForUser(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) ([]models.PartialUser, error) {
	query_string := `
	SELECT
		u.id,
		u.user_handle,
		u.username
	FROM blabber.user u
	INNER JOIN blabber.user_follow uf
		ON uf.followed_id = @id
		AND uf.follower_id = u.id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	row, err := pool.Query(context, query_string, query_args)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	users, err := pgx.CollectRows(row, pgx.RowToStructByName[models.PartialUser])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RetrieveFollowRecordsForUser(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) ([]models.PartialUser, error) {
	query_string := `
	SELECT
		u.id,
		u.user_handle,
		u.username
	FROM blabber.user u
	INNER JOIN blabber.user_follow uf
		ON uf.follower_id = @id
		AND uf.followed_id = u.id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	row, err := pool.Query(context, query_string, query_args)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	users, err := pgx.CollectRows(row, pgx.RowToStructByName[models.PartialUser])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RetrieveBlockRecordsForUser(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) ([]models.PartialUser, error) {
	query_string := `
	SELECT
		u.id,
		u.user_handle,
		u.username
	FROM blabber.user u
	INNER JOIN blabber.user_block uf
		ON uf.blocker_id = @id
		AND uf.blocked_id = u.id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	row, err := pool.Query(context, query_string, query_args)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	users, err := pgx.CollectRows(row, pgx.RowToStructByName[models.PartialUser])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func RetrieveMuteRecordsForUser(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.IndividualUserRequest,
) ([]models.PartialUser, error) {
	query_string := `
	SELECT
		u.id,
		u.user_handle,
		u.username
	FROM blabber.user u
	INNER JOIN blabber.user_mute uf
		ON uf.muter_id = @id
		AND uf.muted_id = u.id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}

	row, err := pool.Query(context, query_string, query_args)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	users, err := pgx.CollectRows(row, pgx.RowToStructByName[models.PartialUser])
	if err != nil {
		return nil, err
	}

	return users, nil
}

func CreateNewFollowingRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.FollowersRequest,
) error {
	// Create following record
	create_record_string := `
	INSERT INTO blabber.user_follow (
		follower_id, followed_id, created_at
	) VALUES (
		@follower, @followed, @created_at
	);
	`
	create_record_args := pgx.NamedArgs{
		"follower":   input.FollowerID,
		"followed":   input.FollowedID,
		"created_at": time.Now(),
	}
	_, err := pool.Exec(context, create_record_string, create_record_args)
	if err != nil {
		return err
	}

	// Increment follower count for followed_id
	increment_followers_string := `
	UPDATE blabber.user
	SET followers = followers + 1
	WHERE id = @followed;
	`
	increment_followers_args := pgx.NamedArgs{
		"followed": input.FollowedID,
	}
	_, err = pool.Exec(context, increment_followers_string, increment_followers_args)
	if err != nil {
		return err
	}

	// Increment follow count for follower_id
	increment_follows_string := `
	UPDATE blabber.user
	SET follows = follows + 1
	WHERE id = @follower;
	`
	increment_follows_args := pgx.NamedArgs{
		"follower": input.FollowerID,
	}
	_, err = pool.Exec(context, increment_follows_string, increment_follows_args)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFollowingRecord(
	context *gin.Context,
	pool *pgxpool.Pool,
	input models.FollowersRequest,
) error {
	// Delete following record
	delete_record_string := `
	DELETE FROM blabber.user_follow
	WHERE follower_id = @follower AND followed_id = @followed;
	`
	delete_record_args := pgx.NamedArgs{
		"follower":   input.FollowerID,
		"followed":   input.FollowedID,
		"created_at": time.Now(),
	}
	_, err := pool.Exec(context, delete_record_string, delete_record_args)
	if err != nil {
		return err
	}

	// Decrement follower count for followed_id
	decrement_followers_string := `
	UPDATE blabber.user
	SET followers = GREATEST(followers - 1, 0)
	WHERE id = @followed;
	`
	decrement_followers_args := pgx.NamedArgs{
		"followed": input.FollowedID,
	}
	_, err = pool.Exec(context, decrement_followers_string, decrement_followers_args)
	if err != nil {
		return err
	}

	// Decrement follow count for follower_id
	decrement_follows_string := `
	UPDATE blabber.user
	SET follows = GREATEST(follows - 1, 0)
	WHERE id = @follower;
	`
	decrement_follows_args := pgx.NamedArgs{
		"follower": input.FollowerID,
	}
	_, err = pool.Exec(context, decrement_follows_string, decrement_follows_args)
	if err != nil {
		return err
	}

	return nil
}
