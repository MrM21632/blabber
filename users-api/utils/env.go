package utils

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	SnowflakeEpoch    uint64 = 1288834974657
	DefaultServerID   uint64 = 0
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

// GetUidgenEpoch retrieves the user UID generation epoch from environment.
//
// Epoch is configurable using the UIDGEN_EPOCH_MS environment variable. This allows for quick,
// reconfigurable deployemnts, especially when scale is required.
//
// If the epoch is not found, we default to the epoch used by Twitter's original Snowflake
// algorithm.
func GetUidgenEpoch() uint64 {
	var epoch uint64
	var err error
	if epoch, err = GetenvInteger("UIDGEN_EPOCH_MS"); err != nil {
		log.Warn("get uidgen epoch failed: envvar UIDGEN_EPOCH_MS not found, using default")
		return SnowflakeEpoch
	}

	return epoch
}

// GetUidgenServerID retrieves the user UID generation server ID from environment.
//
// Server ID is configurable using the UIDGEN_NODE_ID environment variable. This allows for quick,
// reconfigurable deployments, especially when scale is required.
//
// If the server ID is not found, we default to 0.
func GetUidgenServerID() uint64 {
	var server_id uint64
	var err error
	if server_id, err = GetenvInteger("UIDGEN_NODE_ID"); err != nil {
		log.Warn("get uidgen server id failed: envvar UIDGEN_NODE_ID not found, using default")
		return DefaultServerID
	}

	return server_id
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
