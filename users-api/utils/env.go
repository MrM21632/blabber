package utils

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultServerPort uint64 = 8080
)

func getenvStr(key string) (string, error) {
	result := os.Getenv(key)
	if result == "" {
		return result, fmt.Errorf(
			"getenv: Environment variable %s is empty or missing",
			key,
		)
	}

	return result, nil
}

// GetenvInteger retrieves an integer value from environment.
// Returns an error if the environment variable is not found.
func GetenvInteger(key string) (uint64, error) {
	str, err := getenvStr(key)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}

// GetServerPort retrieves the main server port from environment.
//
// Server port is configurable using the SERVER_PORT environment variable. This allows for quick,
// reconfigurable deployments, especially when scale is required.
//
// If the server port is not found, we default to 8080.
func GetServerPort() uint64 {
	var server_port uint64
	var err error
	if server_port, err = GetenvInteger("SERVER_PORT"); err != nil {
		log.Warn("get server port failed: envvar SERVER_PORT not found, using default")
		return DefaultServerPort
	}

	return server_port
}

// GetDatabaseURL retrieves the database connection URL from environment.
//
// Database URL is configurable using the DATABASE_URL environment variable. This allows for quick,
// reconfigurable deployments, especially when scale is required.
//
// If the URL is not found, the program panics.
func GetDatabaseURL() string {
	var database_url string
	var err error
	if database_url, err = getenvStr("DATABASE_URL"); err != nil {
		log.Error("get database url failed: envvar DATABASE_URL not found, cannot continue")
		panic("failed to get database connection details")
	}

	return database_url
}
