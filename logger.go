package golog

import (
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
	logger.Appenders = append(logger.Appenders, newBasicAppender())

	return logger
}

// AddAppender adds one or more appenders to the logger.
func (logger *Logger) AddAppender(appender ...Appender) {
	logger.Appenders = append(logger.Appenders, appender...)
}

// newEntry gets a new entry from the pool.
func (logger *Logger) newEntry() *Entry {
	entry := logger.entries.Get().(*Entry)
	entry.logger = logger

	return entry
}

// freeEntry release the entry into the pool.
func (logger *Logger) freeEntry(entry *Entry) {
	logger.entries.Put(entry)
}

func (logger *Logger) Log(level Level, message string) {
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Log(level, message)
		logger.freeEntry(entry)
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
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Logf(level, message, args...)
		logger.freeEntry(entry)
	}
}

func (logger *Logger) Panicf(message string, args ...interface{}) {
	logger.Logf(PanicLevel, message, args...)
}

func (logger *Logger) Fatalf(message string, args ...interface{}) {
	logger.Logf(FatalLevel, message, args...)
}

func (logger *Logger) Errorf(message string, args ...interface{}) {
	logger.Logf(ErrorLevel, message, args...)
}

func (logger *Logger) Warnf(message string, args ...interface{}) {
	logger.Logf(WarnLevel, message, args...)
}

func (logger *Logger) Infof(message string, args ...interface{}) {
	logger.Logf(InfoLevel, message, args...)
}

func (logger *Logger) Debugf(message string, args ...interface{}) {
	logger.Logf(DebugLevel, message, args...)
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Logln(level, args...)
		logger.freeEntry(entry)
	}
}

func (logger *Logger) Panicln(args ...interface{}) {
	logger.Logln(PanicLevel, args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	logger.Logln(FatalLevel, args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	logger.Logln(ErrorLevel, args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	logger.Logln(WarnLevel, args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	logger.Logln(InfoLevel, args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	logger.Logln(DebugLevel, args...)
}

func (logger *Logger) isLevelEnabled(level Level) bool {
	return logger.Level >= level
}
