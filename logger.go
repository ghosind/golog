package golog

import (
	"sync"
)

type Logger struct {
	Level       Level
	Appenders   []Appender
	LevelLabels map[Level]string

	entries *sync.Pool
}

// New creates and returns a new logger.
func New() *Logger {
	return &Logger{
		Level:     InfoLevel,
		Appenders: make([]Appender, 0),
		entries: &sync.Pool{
			New: func() any {
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

// GetLevelLabel returns the string representation of the specific level.
func (logger *Logger) GetLevelLabel(level Level) string {
	label := ""
	if logger.LevelLabels != nil {
		label = logger.LevelLabels[level]
	}

	if label == "" {
		label = defaultLevelLabels[level]
	}

	return label
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

func (logger *Logger) Print(message string) {
	logger.Log(InfoLevel, message)
}

func (logger *Logger) Logf(level Level, message string, args ...any) {
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Logf(level, message, args...)
		logger.freeEntry(entry)
	}
}

func (logger *Logger) Panicf(message string, args ...any) {
	logger.Logf(PanicLevel, message, args...)
}

func (logger *Logger) Fatalf(message string, args ...any) {
	logger.Logf(FatalLevel, message, args...)
}

func (logger *Logger) Errorf(message string, args ...any) {
	logger.Logf(ErrorLevel, message, args...)
}

func (logger *Logger) Warnf(message string, args ...any) {
	logger.Logf(WarnLevel, message, args...)
}

func (logger *Logger) Infof(message string, args ...any) {
	logger.Logf(InfoLevel, message, args...)
}

func (logger *Logger) Debugf(message string, args ...any) {
	logger.Logf(DebugLevel, message, args...)
}

func (logger *Logger) Printf(message string, args ...any) {
	logger.Logf(InfoLevel, message, args...)
}

func (logger *Logger) Logln(level Level, args ...any) {
	if logger.isLevelEnabled(level) {
		entry := logger.newEntry()
		entry.Logln(level, args...)
		logger.freeEntry(entry)
	}
}

func (logger *Logger) Panicln(args ...any) {
	logger.Logln(PanicLevel, args...)
}

func (logger *Logger) Fatalln(args ...any) {
	logger.Logln(FatalLevel, args...)
}

func (logger *Logger) Errorln(args ...any) {
	logger.Logln(ErrorLevel, args...)
}

func (logger *Logger) Warnln(args ...any) {
	logger.Logln(WarnLevel, args...)
}

func (logger *Logger) Infoln(args ...any) {
	logger.Logln(InfoLevel, args...)
}

func (logger *Logger) Debugln(args ...any) {
	logger.Logln(DebugLevel, args...)
}

func (logger *Logger) Println(args ...any) {
	logger.Logln(InfoLevel, args...)
}

func (logger *Logger) isLevelEnabled(level Level) bool {
	return logger.Level >= level
}
