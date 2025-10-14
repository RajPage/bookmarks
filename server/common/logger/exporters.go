package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Exporter interface {
	Export(entry LogEntry) error
}

type ConsoleExporter struct {
	Zlog zerolog.Logger
}

type ExporterFlavor string

const (
	FlavorConsole ExporterFlavor = "console"
)

func NewExporter(flavor ExporterFlavor, level Level) Exporter {
	switch flavor {
	case FlavorConsole:
		return NewConsoleExporter(level)
	}
	return nil
}

func NewConsoleExporter(level Level) *ConsoleExporter {
	var output io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.DateTime,
	}
	zlog := zerolog.New(output).
		Level(zerolog.Level(level)).
		With().
		Timestamp().
		// Str("git_revision", gitRevision).
		// Str("go_version", buildInfo.GoVersion).
		Logger()
	return &ConsoleExporter{Zlog: zlog}
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
