package logger

import "context"

type Logger struct {
	processor *Processor
}

// TODO: infer context.

func (l *Logger) Debug(ctx context.Context, msg string) {
	l.Log(ctx, DEBUG, msg)
}

func (l *Logger) Info(ctx context.Context, msg string) {
	l.Log(ctx, INFO, msg)
}

func (l *Logger) Warn(ctx context.Context, msg string) {
	l.Log(ctx, WARN, msg)
}

func (l *Logger) Error(ctx context.Context, msg string) {
	l.Log(ctx, ERROR, msg)
}

func (l *Logger) Log(ctx context.Context, level Level, msg string) {
	entry := l.processor.Process(ctx, level, msg)
	l.processor.Export(entry)
}

func (l *Logger) GetLogLevel() Level {
	return l.processor.minLevel
}
