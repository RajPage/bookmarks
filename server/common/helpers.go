package common

import (
	"bookmarks/server/common/logger"
	"bookmarks/server/config/environment"
	"context"
	"os"
)

func GetEnvValue(key string) string {
	defaultValue, exists := environment.DefaultEnvVars[key]
	if !exists {
		// I just realized this can cause a cyclic dependency if logger is not initialized yet
		// For now moving this inside the if block to avoid the cycle
		log := logger.GetLogger()
		log.Warn(context.TODO(), "No default value found for env var key: "+key)
		return ""
	}
	returnValue := defaultValue

	val, exists := os.LookupEnv(key)
	if exists {
		returnValue = val
	}

	return returnValue
}
