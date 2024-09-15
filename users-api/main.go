package main

import (
	"fmt"
	"net/http"
	"users-api/api"
	"users-api/middleware"
	"users-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	godotenv.Load()

	server_port := utils.GetServerPort()
	PasswordParams := utils.GetArgon2IDConfig()
	UIDGenerator, err := utils.InitializeNode()
	if err != nil {
		log.Error("encountered error while initializing uidgen node: " + err.Error())
		return
	}
	// TODO: actually name this variable when used
	_ = api.UserServer{
		UIDGenerator: UIDGenerator,
		Argon2Params: &PasswordParams,
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
