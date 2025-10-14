package environment

import "bookmarks/server/common/constants"

const (
	EnvServerEnvironment = "BOOKMARKS_SERVER_ENVIRONMENT"
	EnvServerLogLevel    = "BOOKMARKS_SERVER_LOG_LEVEL"
)

var DefaultEnvVars = map[string]string{
	EnvServerEnvironment: constants.ServerEnvironmentDevelopment,
	EnvServerLogLevel:    constants.ServerLogLevel,
}
