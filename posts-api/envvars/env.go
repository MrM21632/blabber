package envvars

import (
	"fmt"
	"os"
	"strconv"
)

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

func GetenvBoolean(key string) (bool, error) {
	str, err := GetenvStr(key)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}

	return result, nil
}
