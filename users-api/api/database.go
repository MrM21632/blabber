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
