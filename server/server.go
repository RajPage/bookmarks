package server

import (
	"bookmarks/server/common/logger"
	"bookmarks/server/config"
	"context"
	"strconv"

	"github.com/rs/zerolog"
)

func initServer() {
	config.InitializeConfig()
	startLogger()
	// context
	// db
	// Set keys
	// Middleware
	// routes
	// swagger (for dev)
}

func startLogger() {
	logLevel := config.GetConfig().LogLevel
	level, err := strconv.Atoi(logLevel)
	if err != nil {
		level = int(zerolog.InfoLevel)
	}
	consoleExporter := logger.NewExporter(logger.FlavorConsole, logger.Level(level))
	logger.InitializeLogger(logger.Config{
		MinLevel:  logger.Level(level),
		Exporters: []logger.Exporter{consoleExporter},
	})
}

func Start() {
	ctx := context.Background()

	initServer()

	log := logger.GetLogger()
	log.Info(ctx, "Starting server...")
}
