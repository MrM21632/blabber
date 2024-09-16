package api

import (
	"net/http"
	"time"
	"users-api/models"
	"users-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServer struct {
	Argon2Params *utils.Argon2IDParams
	DatabasePool *pgxpool.Pool
}

// POST /users
func (u UserServer) CreateUser(context *gin.Context) {
	var err error

	var input models.CreateUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var new_user_id uuid.UUID
	if new_user_id, err = uuid.NewV7(); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error occurred creating new user id: " + err.Error()},
		)
		return
	}

	var password_hash string
	if password_hash, err = utils.GenerateHash(input.Password, u.Argon2Params); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error occurred hashing password: " + err.Error()},
		)
		return
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
		"id":        new_user_id,
		"username":  input.Username,
		"handle":    input.Handle,
		"bio":       *input.Bio,
		"email":     input.Email,
		"hash":      password_hash,
		"createdAt": time.Now(),
		"updatedAt": time.Now(),
		"followers": 0,
		"follows":   0,
	}

	_, err = u.DatabasePool.Exec(context, query_string, query_args)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error occurred writing new user record: " + err.Error()},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"user_id": new_user_id},
	)
}

// GET /users
func (u UserServer) GetUser(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query_string := `
	SELECT * FROM blabber.user
	WHERE id = @id;
	`
	query_args := pgx.NamedArgs{
		"id": input.ID,
	}
	row, err := u.DatabasePool.Query(context, query_string, query_args)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error occurred retrieving user record: " + err.Error()},
		)
		return
	}
	defer row.Close()

	user, err := pgx.CollectOneRow(row, pgx.RowToStructByName[models.User])
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "error occurred retrieving user record: " + err.Error()},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"user": user},
	)
}

// GET /users/followers
func (u UserServer) GetFollowers(context *gin.Context) {}

// GET /users/follows
func (u UserServer) GetFollows(context *gin.Context) {}

// GET /users/blocks
func (u UserServer) GetBlocks(context *gin.Context) {}

// GET /users/mutes
func (u UserServer) GetMutes(context *gin.Context) {}

// PATCH /users
func (u UserServer) UpdateUser(context *gin.Context) {}

// PATCH /users/password
func (u UserServer) UpdatePassword(context *gin.Context) {}

// DELETE /users
func (u UserServer) DeleteUser(context *gin.Context) {}

// POST /users/follow
func (u UserServer) FollowUser(context *gin.Context) {}

// DELETE /users/follow
func (u UserServer) UnfollowUser(context *gin.Context) {}

// POST /users/block
func (u UserServer) BlockUser(context *gin.Context) {}

// DELETE /users/block
func (u UserServer) UnblockUser(context *gin.Context) {}

// POST /users/mute
func (u UserServer) MuteUser(context *gin.Context) {}

// DELETE /users/mute
func (u UserServer) UnmuteUser(context *gin.Context) {}
