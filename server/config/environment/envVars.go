package environment

import "bookmarks/server/common/constants"

const (
	EnvServerEnvironment = "BOOKMARKS_SERVER_ENVIRONMENT"
	EnvServerLogLevel    = "BOOKMARKS_SERVER_LOG_LEVEL"
	EnvDbDsn             = "BOOKMARKS_DB_DSN"
)

var DefaultEnvVars = map[string]string{
	EnvServerEnvironment: constants.ServerEnvironmentDevelopment,
	EnvServerLogLevel:    constants.ServerLogLevel,
	EnvDbDsn:             constants.DbDsn,
}
