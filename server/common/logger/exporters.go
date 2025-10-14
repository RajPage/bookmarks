package logger

import (
	"github.com/rs/zerolog"
)

type Exporter interface {
	Export(entry LogEntry) error
}

type ConsoleExporter struct {
	Zlog zerolog.Logger
}

func (ce *ConsoleExporter) Export(entry LogEntry) error {
	zlogEvent := ce.Zlog.WithLevel(zerolog.InfoLevel)
	zlogEvent = zlogEvent.Str("message", entry.Message)
	zlogEvent.Send()
	return nil
}

// Placeholders
// type FileExporter struct {}
// type RemoteExporter struct {}
