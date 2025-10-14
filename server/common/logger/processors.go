package logger

import "context"

type Processor struct {
	minLevel  Level
	exporters []Exporter
}

func (cp *Processor) Export(entry *LogEntry) error {
	if entry.Level < cp.minLevel {
		return nil
	}
	for _, exporter := range cp.exporters {
		if err := exporter.Export(*entry); err != nil {
			return err
		}
	}
	return nil
}

func (cp *Processor) Process(ctx context.Context, level Level, msg string) *LogEntry {
	entry := &LogEntry{
		Level:   level,
		Message: msg,
	}
	return entry
}
