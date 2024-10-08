package utils

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultServerPort uint64 = 8080
	DefaultLifespan   uint64 = 1
)

// GetenvStr retrieves a string value from environment.
// Returns an error if the environment variable is not found.
func GetenvStr(key string) (string, error) {
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
	str, err := GetenvStr(key)
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
	if database_url, err = GetenvStr("DATABASE_URL"); err != nil {
		log.Error("get database url failed: envvar DATABASE_URL not found, cannot continue")
		panic("failed to get database connection details")
	}

	return database_url
}

// GetJWTSecret retrieves the JWT signing secret key from environment.
//
// JWT secret key is configurable using the DATABASE_URL environment variable. This allows for quick,
// reconfigurable deployments, especially when scale is required.
//
// If the secret is not found, the program panics.
func GetJWTSecret() string {
	var jwt_secret string
	var err error
	if jwt_secret, err = GetenvStr("JWT_SECRET"); err != nil {
		log.Error("get JWT secret failed: envvar JWT_SECRET not found, cannot continue")
		panic("failed to get JWT secret")
	}

	return jwt_secret
}

// GetJWTLifespan retrieves the JWT lifespan, in hours, from environment.
//
// JWT lifespan is configurable using the JWT_LIFESPAN_HOURS environment variable. This allows for
// quick, reconfigurable deployments, especially when scale is required.
//
// If the lifespan is not found, we default to 1 hour.
func GetJWTLifespan() uint64 {
	var lifespan uint64
	var err error
	if lifespan, err = GetenvInteger("JWT_LIFESPAN_HOURS"); err != nil {
		log.Warn("get JWT lifespan failed: envvar JWT_LIFESPAN_HOURS not found, using default")
		return DefaultLifespan
	}

	return lifespan
}
