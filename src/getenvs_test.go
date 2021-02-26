package getenvs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	// SET_STRING_GETENV must have been set before
	value := GetEnvString("SET_STRING_GETENV", "default-string-value")
	value2 := GetEnvString("NOT_SET_STRING_GETENV", "default-string-value")

	assert.Equal(t, value, "env-has-set")
	assert.Equal(t, value2, "default-string-value")
}
