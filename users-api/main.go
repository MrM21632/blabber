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

	uid_generator, err := utils.InitializeNode()
	if err != nil {
		log.Error("encountered error while initializing uidgen node: " + err.Error())
		return
	}

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

	// TODO: actually name this variable when used
	_ = api.UserServer{
		UIDGenerator: uid_generator,
		Argon2Params: &password_params,
		DatabasePool: pool,
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.SetTrustedProxies(nil)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "ENDPOINT_NOT_FOUND", "message": "endpoint not found"})
	})

	r.Run(fmt.Sprintf(":%d", server_port))
}
