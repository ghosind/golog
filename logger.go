package golog

import (
	"fmt"
	"sync"
)

type Logger struct {
	Level     Level
	Appenders []Appender

	entries *sync.Pool
}

// New creates and returns a new logger.
func New() *Logger {
	return &Logger{
		Level:     InfoLevel,
		Appenders: make([]Appender, 0),
		entries: &sync.Pool{
			New: func() interface{} {
				return newEntry()
			},
		},
	}
}

// Default creates a new logger with default settings.
func Default() *Logger {
	logger := New()
	logger.Appenders = append(logger.Appenders, NewConsoleAppender(ConsoleAppenderConfig{}))

	return logger
}

func (logger *Logger) newEntry() *Entry {
	entry := logger.entries.Get().(*Entry)

	return entry
}

func (logger *Logger) finalizeEntry(entry *Entry) {
	logger.entries.Put(entry)
}

func (logger *Logger) Log(level Level, message string) {
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Log(level, message)
		logger.finalizeEntry(entry)
	}
}

func (logger *Logger) Panic(message string) {
	logger.Log(PanicLevel, message)
}

func (logger *Logger) Fatal(message string) {
	logger.Log(FatalLevel, message)
}

func (logger *Logger) Error(message string) {
	logger.Log(ErrorLevel, message)
}

func (logger *Logger) Warn(message string) {
	logger.Log(WarnLevel, message)
}

func (logger *Logger) Info(message string) {
	logger.Log(InfoLevel, message)
}

func (logger *Logger) Debug(message string) {
	logger.Log(DebugLevel, message)
}

func (logger *Logger) Logf(level Level, message string, args ...interface{}) {
	logger.Log(level, fmt.Sprintf(message, args...))
}

func (logger *Logger) Panicf(message string, args ...interface{}) {
	logger.Logf(PanicLevel, message)
}

func (logger *Logger) Fatalf(message string, args ...interface{}) {
	logger.Logf(FatalLevel, message)
}

func (logger *Logger) Errorf(message string, args ...interface{}) {
	logger.Logf(ErrorLevel, message)
}

func (logger *Logger) Warnf(message string, args ...interface{}) {
	logger.Logf(WarnLevel, message)
}

func (logger *Logger) Infof(message string, args ...interface{}) {
	logger.Logf(InfoLevel, message)
}

func (logger *Logger) Debugf(message string, args ...interface{}) {
	logger.Logf(DebugLevel, message)
}

func (logger *Logger) isLevelEnabled(level Level) bool {
	return logger.Level >= level
}
