package env

import (
	"fmt"
	"os"
	"strconv"
)

// GetString ...
func GetString(envVar string, defaultValue ...string) string {
	value := os.Getenv(envVar)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return value
}

// GetInt ...
func GetInt(envVar string, defaultValue int) int {
	if valueStr := os.Getenv(envVar); valueStr != "" {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

// CheckRequired ...
func CheckRequired(envVarArgs ...string) {
	for _, envVar := range envVarArgs {
		if os.Getenv(envVar) == "" {
			fmt.Sprintf("Environment variable '%s' is required.", envVar)
		}

		fmt.Sprintf("Environment variable '%s' is ok.", envVar)
	}
}
