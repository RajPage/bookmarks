package server

import (
	"bookmarks/server/common/logger"
	"context"
	"os"

	"github.com/rs/zerolog"
)

func Start() {
	ctx := context.TODO()

	consoleExporter := &logger.ConsoleExporter{
		Zlog: zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger(),
	}

	logger.InitializeLogger(logger.Config{
		MinLevel:  logger.DEBUG,
		Exporters: []logger.Exporter{consoleExporter},
	})

	// config
	// logger
	// context
	// db
	// check env vars
	// Set keys
	// Middleware
	// routes
	// swagger (for dev)
	log := logger.GetLogger()
	log.Info(ctx, "Starting server...")
	// start server
}
