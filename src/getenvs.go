package getenvs

import "os"

// GetEnvString gets the environment variable for a key and if that env-var hasn't been set it returns the default value
func GetEnvString(key string, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultVal
	}
	return value
}
