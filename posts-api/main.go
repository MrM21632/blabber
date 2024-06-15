package main

import (
	"fmt"
	"net/http"
	"posts/envvars"
	"posts/middleware"

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

	server_port, err := envvars.GetenvInteger("SERVER_PORT")
	if err != nil {
		fmt.Printf("Encountered error when retrieving server port, setting to default: %s", err.Error())
		server_port = 8081
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggingMiddleware())
	r.SetTrustedProxies(nil)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "ENDPOINT_NOT_FOUND", "message": "Endpoint not found"})
	})

	r.Run(fmt.Sprintf(":%d", server_port))
}
