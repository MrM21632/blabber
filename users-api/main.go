package main

import (
	"context"
	"fmt"
	"net/http"
	"users-api/api"
	"users-api/middleware"
	"users-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	pgx_uuid "github.com/vgarvardt/pgx-google-uuid/v5"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	godotenv.Load()

	server_port := utils.GetServerPort()

	password_params := utils.GetArgon2IDConfig()

	// Be sure to include ?search_path=blabber in the database URL
	config, err := pgxpool.ParseConfig(utils.GetDatabaseURL())
	if err != nil {
		log.Error("encountered error while parsing config string: ", err.Error())
		return
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgx_uuid.Register(conn.TypeMap())
		return nil
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Error("encountered error while creating connection pool: " + err.Error())
		return
	}
	defer pool.Close()

	server := api.UserServer{
		Argon2Params: &password_params,
		DatabasePool: pool,
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.SetTrustedProxies(nil)

	// GET endpoints
	r.GET("/users", server.GetUser)
	r.GET("/users/followers", server.GetFollowers)
	r.GET("/users/follows", server.GetFollows)
	r.GET("/users/blocks", server.GetBlocks)
	r.GET("/users/mutes", server.GetMutes)

	// POST endpoints
	r.POST("/users", server.CreateUser)
	r.POST("/users/follow", server.FollowUser)
	r.POST("/users/block", server.BlockUser)
	r.POST("/users/mute", server.MuteUser)

	// PUT and PATCH endpoints
	r.PATCH("/users", server.UpdateUser)
	r.PATCH("/users/password", server.UpdatePassword)

	// DELETE endpoints
	r.DELETE("/users", server.DeleteUser)
	r.DELETE("/users/follow", server.UnfollowUser)
	r.DELETE("/users/block", server.UnblockUser)
	r.DELETE("/users/mute", server.UnmuteUser)

	// Invalid routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "ENDPOINT_NOT_FOUND", "message": "endpoint not found"})
	})

	r.Run(fmt.Sprintf(":%d", server_port))
}
