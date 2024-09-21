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

	_, err = WriteNewUserRecord(context, u.DatabasePool, input, password_hash, new_user_id)
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			context.JSON(
				http.StatusConflict,
				gin.H{"error": fmt.Sprintf("received violative entity: %s", err.Error())},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := RetrieveUserRecord(context, u.DatabasePool, input)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			context.JSON(
				http.StatusNotFound,
				gin.H{"error": fmt.Sprintf("user entity with id=%s not found: %s", input.ID.String(), err.Error())},
			)
		} else {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := RetrieveFollowerRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := RetrieveFollowRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := RetrieveBlockRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
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
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := RetrieveMuteRecordsForUser(context, u.DatabasePool, input)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("unexpected error occurred: %s", err.Error())},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"mutes": users})
}

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
