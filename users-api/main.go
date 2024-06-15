package main

import (
	"fmt"
	"net/http"
	"users/envvars"
	"users/middleware"
	"users/uidgen"
	"users/users"

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
	users.ConnectToDatabase()

	UidGenNode, err := uidgen.InitializeNode()
	if err != nil {
		fmt.Printf("Encountered error while initializing uidgen node: %s", err.Error())
		return
	}
	PasswordParams := &users.Argon2idParams{
		Memory:  64 * 1024, // 64 MiB
		Time:    3,
		Threads: 2,
		Saltlen: 16,
		Hashlen: 32,
	}
	server := users.UserServer{
		UidGenNode: UidGenNode,
		PassParams: PasswordParams,
	}

	server_port, err := envvars.GetenvInteger("SERVER_PORT")
	if err != nil {
		fmt.Printf("Encountered error when retrieving server port, setting to default: %s", err.Error())
		server_port = 8080
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.SetTrustedProxies(nil)

	r.POST("/users", server.CreateUser)
	r.GET("/users", server.GetUser)
	r.GET("/users/followers", server.GetFollowers)
	r.GET("/users/follows", server.GetFollows)
	r.PATCH("/users", server.UpdateUser)
	r.PATCH("/users/password", server.UpdatePassword)
	r.DELETE("/users", server.DeleteUser)

	r.POST("/follow", server.FollowUser)
	r.DELETE("/follow", server.UnfollowUser)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "ENDPOINT_NOT_FOUND", "message": "Endpoint not found"})
	})

	r.Run(fmt.Sprintf(":%d", server_port))
}
