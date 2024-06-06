package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"users/envvars"
	"users/uidgen"
	"users/users"
	"users/users/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	UidGenNode     *uidgen.UniqueIdGenerator
	PasswordParams *users.Argon2idParams
)

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

	server_port, err := envvars.GetenvInteger("SERVER_PORT")
	if err != nil {
		fmt.Printf("Encountered error when retrieving server port, setting to default: %s", err.Error())
		server_port = 8080
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/users", func(c *gin.Context) {
		var input models.CreateUserRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		new_user_id := UidGenNode.GeanerateId()
		result, err := users.WriteNewUserRecord(
			strconv.FormatUint(uint64(new_user_id), 10),
			&input,
			PasswordParams,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"user": *result})
	})

	r.GET("/users", func(c *gin.Context) {
		var input models.IndividualUserRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := users.GetUserRecord(input.ID)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": *result})
	})

	r.PATCH("/users", func(c *gin.Context) {
		var input models.UpdateUserRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := users.UpdateUserRecord(&input); err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.PATCH("/users/password", func(c *gin.Context) {
		var input models.UpdateUserPasswordRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		curr_password, err := users.GetUserPasswordHash(input.ID)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		match, err := users.ComparePasswordToHash(input.OldPassword, *curr_password)
		if !match {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}

		new_hash, err := users.GenerateHash(input.NewPassword, PasswordParams)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		if err := users.UpdatePasswordHashForUser(input.ID, new_hash); err != nil {
			if strings.Contains(err.Error(), "record not found") {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Endpoint not found"})
	})

	r.Run(fmt.Sprintf(":%d", server_port))
}
