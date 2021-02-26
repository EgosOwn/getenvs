package getenvs

import (
	"os"
	"strconv"
)

// GetEnvString gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultVal
	}
	return value
}

// GetEnvBool gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvBool(key string, defaultVal bool) (bool, error) {
	strvalue := os.Getenv(key)
	if len(strvalue) == 0 {
		value := defaultVal
		return value, nil
	}

	value, err := strconv.ParseBool(strvalue)
	return value, err
}
