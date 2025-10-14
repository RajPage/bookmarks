package logger

import "sync"

type Config struct {
	MinLevel  Level
	Exporters []Exporter
}

var (
	singletonLogger *Logger
	once            sync.Once
)

func InitializeLogger(cfg Config) *Logger {
	once.Do(func() {
		processor := &Processor{
			minLevel:  cfg.MinLevel,
			exporters: cfg.Exporters,
		}
		singletonLogger = &Logger{
			processor: processor,
		}
	})
	return singletonLogger
}

func GetLogger() *Logger {
	if singletonLogger == nil {
		panic("Logger not initialized. Call InitializeLogger first.")
	}
	return singletonLogger
}
