package server

import (
	"bookmarks/server/common/logger"
	"bookmarks/server/config"
	"bookmarks/server/internal/model"
	"context"
	"strconv"

	"github.com/rs/zerolog"
)

// initServer initializes the server by setting up the magic
func initServer() {
	config.InitializeConfig()
	startLogger()

	dsn := config.GetConfig().DbDsn
	model.InitializeDatabase(dsn)

	// context
	// Set keys
	// Middleware
	// routes
	// swagger (for dev)
}

// startLogger initializes the logger based on the configuration.
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

// Start starts the server.
func Start() {
	ctx := context.Background()

	initServer()

	log := logger.GetLogger()
	log.Info(ctx, "Starting server...")
}
