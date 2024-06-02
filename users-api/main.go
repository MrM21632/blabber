package main

import (
	"fmt"
	"users/envvars"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// uidgen_node, err := uidgen.InitializeNode()
	// if err != nil {
	// 	fmt.Printf("Encountered error while initializing uidgen node: %s", err.Error())
	// 	return
	// }

	server_port, err := envvars.GetenvInteger("SERVER_PORT")
	if err != nil {
		fmt.Printf("Encountered error when retrieving server port, setting to default: %s", err.Error())
		server_port = 8080
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Run(fmt.Sprintf(":%d", server_port))
}
