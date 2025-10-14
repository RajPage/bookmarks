package server

import (
	"bookmarks/server/common"
	"bookmarks/server/common/logger"
	"context"
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

func startLogger() {
	logLevel := common.GetEnvValue("BOOKMARKS_SERVER_LOG_LEVEL")
	level, err := strconv.Atoi(logLevel)
	if err != nil {
		level = int(zerolog.InfoLevel)
	}
	consoleExporter := &logger.ConsoleExporter{
		Zlog: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger(),
	}
	logger.InitializeLogger(logger.Config{
		MinLevel:  logger.Level(level),
		Exporters: []logger.Exporter{consoleExporter},
	})
}

func Start() {
	ctx := context.Background()
	startLogger()

	// config
	// context
	// db
	// Set keys
	// Middleware
	// routes
	// swagger (for dev)
	log := logger.GetLogger()
	log.Info(ctx, "Starting server...")
	// start server
}
