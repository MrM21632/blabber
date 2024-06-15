package users

import (
	"net/http"
	"strconv"
	"strings"
	"users/uidgen"
	"users/users/models"

	"github.com/gin-gonic/gin"
)

type UserServer struct {
	UidGenNode *uidgen.UniqueIdGenerator
	PassParams *Argon2idParams
}

// POST /users
func (u UserServer) CreateUser(context *gin.Context) {
	var input models.CreateUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_user_id := u.UidGenNode.GeanerateId()
	result, err := WriteNewUserRecord(
		strconv.FormatUint(uint64(new_user_id), 10),
		&input,
		u.PassParams,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": *result})
}

// POST /follow
func (u UserServer) FollowUser(context *gin.Context) {
	var input models.FollowUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := WriteNewFollowingRecord(&input)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"following": *result})
}

// GET /users (individual accounts only)
func (u UserServer) GetUser(context *gin.Context) {
	var input models.IndividualUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := GetUserRecord(input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"user": *result})
}

// GET /users/followers
func (u UserServer) GetFollowers(context *gin.Context) {
	var input models.GetFollowersRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := GetFollowersForUser(input.FollowedID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"followers": result})
}

// GET /users/follows
func (u UserServer) GetFollows(context *gin.Context) {
	var input models.GetFollowsRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := GetFollowsForUser(input.FollowerID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"follows": result})
}

// PATCH /users
func (u UserServer) UpdateUser(context *gin.Context) {
	var input models.UpdateUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateUserRecord(&input); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}

// PATCH /users/password
func (u UserServer) UpdatePassword(context *gin.Context) {
	var input models.UpdateUserPasswordRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curr_password, err := GetUserPasswordHash(input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	match, err := ComparePasswordToHash(input.OldPassword, *curr_password)
	if !match {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
	}

	new_hash, err := GenerateHash(input.NewPassword, u.PassParams)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
	}
	if err := UpdatePasswordHashForUser(input.ID, new_hash); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}

// DELETE /users
func (u UserServer) DeleteUser(context *gin.Context) {
	var input models.IndividualUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DeleteUserRecord(input.ID); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if strings.Contains(err.Error(), "record already deleted") {
			context.JSON(http.StatusGone, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}

// DELETE /follow
func (u UserServer) UnfollowUser(context *gin.Context) {
	var input models.FollowUserRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DeleteFollowingRecord(input.FollowerID, input.FollowedID); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if strings.Contains(err.Error(), "record already deleted") {
			context.JSON(http.StatusGone, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusNoContent, gin.H{})
}
