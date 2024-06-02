package main

import (
	"fmt"
	"users/envvars"
	"users/uidgen"
	"users/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	UidGenNode *uidgen.UniqueIdGenerator
	PassParams *users.Argon2idParams
)

func main() {
	godotenv.Load()
	UidGenNode, err := uidgen.InitializeNode()
	if err != nil {
		fmt.Printf("Encountered error while initializing uidgen node: %s", err.Error())
		return
	}
	_ = UidGenNode // TODO: Remove once we're using this variable

	PassParams := &users.Argon2idParams{
		Memory:  64 * 1024,
		Time:    3,
		Threads: 2,
		Saltlen: 16,
		Hashlen: 32,
	}
	_ = PassParams // TODO: Remove once we're using this variable

	server_port, err := envvars.GetenvInteger("SERVER_PORT")
	if err != nil {
		fmt.Printf("Encountered error when retrieving server port, setting to default: %s", err.Error())
		server_port = 8080
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Run(fmt.Sprintf(":%d", server_port))
}
