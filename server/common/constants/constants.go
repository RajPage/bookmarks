package constants

import (
	"strconv"

	"github.com/rs/zerolog"
)

const (
	ServerEnvironmentDevelopment = "development"
	ServerEnvironmentProduction  = "production"
)

var ServerLogLevel = strconv.Itoa(int(zerolog.DebugLevel))

const (
	DbDsn = "host=localhost user=bookmarks password=bookmarks dbname=bookmarks port=5432 sslmode=disable TimeZone=Asia/Kolkata"
)
