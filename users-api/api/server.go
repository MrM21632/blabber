package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}

	var new_user_id uuid.UUID
	if new_user_id, err = uuid.NewV7(); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "error creating new UID: " + err.Error()},
		)
		return
	}

	var password_hash string
	if password_hash, err = utils.GenerateHash(input.Password, u.Argon2Params); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "error hashing password: " + err.Error()},
		)
		return
	}

	_, err = WriteNewUserRecord(context, u.DatabasePool, input, password_hash, new_user_id)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			context.JSON(
				http.StatusConflict,
				gin.H{"code": "DB_CONFLICT", "error": "received violative entity: " + err.Error()},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
			)
		}
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user_id": new_user_id})
}

// GET /users
func (u UserServer) GetUser(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	user, err := RetrieveUserRecord(context, u.DatabasePool, input)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			context.JSON(
				http.StatusNotFound,
				gin.H{
					"code":  "ENTITY_NOT_FOUND",
					"error": fmt.Sprintf("user entity with id=%s not found: %s", input.ID.UUID.String(), err.Error()),
				},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
			)
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": *user})
}

// GET /users/followers
func (u UserServer) GetFollowers(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	users, err := RetrieveFollowerRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"followers": users})
}

// GET /users/follows
func (u UserServer) GetFollows(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	users, err := RetrieveFollowRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"follows": users})
}

// GET /users/blocks
func (u UserServer) GetBlocks(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	users, err := RetrieveBlockRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"blocks": users})
}

// GET /users/mutes
func (u UserServer) GetMutes(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	users, err := RetrieveMuteRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"mutes": users})
}

// PATCH /users
func (u UserServer) UpdateUser(context *gin.Context) {
	var err error

	var input models.UpdateUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	user, err := RetrieveUserRecord(context, u.DatabasePool, models.IndividualUserRequest{ID: input.ID})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			context.JSON(
				http.StatusNotFound,
				gin.H{
					"code":  "ENTITY_NOT_FOUND",
					"error": fmt.Sprintf("user entity with id=%s not found: %s", input.ID.UUID.String(), err.Error()),
				},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
			)
		}
		return
	}

	if _, err = UpdateUserRecord(context, u.DatabasePool, input, *user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// PATCH /users/password
func (u UserServer) UpdatePassword(context *gin.Context) {
	var err error

	var input models.UpdateUserPasswordRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	curr_hash, err := RetrieveUserPassword(context, u.DatabasePool, input.ID.UUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			context.JSON(
				http.StatusNotFound,
				gin.H{
					"code":  "ENTITY_NOT_FOUND",
					"error": fmt.Sprintf("user entity with id=%s not found: %s", input.ID.UUID.String(), err.Error()),
				},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
			)
		}
		return
	}

	match, err := utils.ComparePasswordToHash(input.OldPassword, *curr_hash)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "unexpected error occurred: " + err.Error()},
		)
		return
	}
	if !match {
		context.JSON(http.StatusUnauthorized, gin.H{"code": "AUTH_ERROR", "error": "invalid password"})
		return
	}

	new_hash, err := utils.GenerateHash(input.NewPassword, u.Argon2Params)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"code": "INTERNAL_ERROR", "error": "error hashing password: " + err.Error()},
		)
		return
	}

	if _, err := UpdateUserPassword(context, u.DatabasePool, input.ID.UUID, new_hash); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// DELETE /users
func (u UserServer) DeleteUser(context *gin.Context) {
	var err error

	var input models.IndividualUserRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.ID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "user_id not provided"})
		return
	}

	if _, err := DeleteUserRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// POST /users/follow
func (u UserServer) FollowUser(context *gin.Context) {
	var err error

	var input models.FollowersRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.FollowedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "followed_id not provided"})
		return
	}
	if !input.FollowerID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "follower_id not provided"})
		return
	}

	if err := CreateNewFollowingRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{})
}

// DELETE /users/follow
func (u UserServer) UnfollowUser(context *gin.Context) {
	var err error

	var input models.FollowersRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.FollowedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "followed_id not provided"})
		return
	}
	if !input.FollowerID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "follower_id not provided"})
		return
	}

	if err := DeleteFollowingRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// POST /users/block
func (u UserServer) BlockUser(context *gin.Context) {
	var err error

	var input models.BlocksRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.BlockedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "blocked_id not provided"})
		return
	}
	if !input.BlockerID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "blocker_id not provided"})
		return
	}

	if err := CreateNewBlockRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// DELETE /users/block
func (u UserServer) UnblockUser(context *gin.Context) {
	var err error

	var input models.BlocksRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.BlockedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "blocked_id not provided"})
		return
	}
	if !input.BlockerID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "blocker_id not provided"})
		return
	}

	if err := DeleteBlockRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// POST /users/mute
func (u UserServer) MuteUser(context *gin.Context) {
	var err error

	var input models.MutesRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.MutedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "muted_id not provided"})
		return
	}
	if !input.MuterID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "muter_id not provided"})
		return
	}

	if err := CreateNewMuteRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}

// DELETE /users/mute
func (u UserServer) UnmuteUser(context *gin.Context) {
	var err error

	var input models.MutesRequest
	if err = context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"code": "BINDING_FAILURE", "error": err.Error()})
		return
	}
	if !input.MutedID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "muted_id not provided"})
		return
	}
	if !input.MuterID.Valid {
		context.JSON(http.StatusBadRequest, gin.H{"code": "MISSING_FIELD", "error": "muter_id not provided"})
		return
	}

	if err := DeleteMuteRecord(context, u.DatabasePool, input); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"code": "DB_ERROR", "error": err.Error()})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{})
}
