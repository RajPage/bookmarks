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
