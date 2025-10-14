package server

import (
	"bookmarks/server/common/logger"
	"bookmarks/server/config"
	"context"
	"os"
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

	initServer()

	log := logger.GetLogger()
	log.Info(ctx, "Starting server...")
}
