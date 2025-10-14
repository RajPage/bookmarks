package logger

import "time"

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

type LogEntry struct {
	Level     Level
	Message   string
	Timestamp time.Time
	// Fields    map[string]interface{}
	// TraceID   string
	// SpanID    string
}
