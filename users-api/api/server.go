package api

import (
	"users-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserServer struct {
	Argon2Params *utils.Argon2IDParams
	DatabasePool *pgxpool.Pool
}

// POST /users
func (u UserServer) CreateUser(context *gin.Context) {}

// GET /users
func (u UserServer) GetUser(context *gin.Context) {}

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
