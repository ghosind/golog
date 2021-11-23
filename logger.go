package logger

import (
	"errors"
	"log"
)

const (
	LogLevelOff int = 0
	LogLevelFatal
	LogLevelError
	LogLevelWarn
	LogLevelInfo
	LogLevelDebug
	LogLevelTrace
	LogLevelAll
)

type Logger struct {
	level int

	logger *log.Logger

	name string

	transports []Appender
}

func New(name string) *Logger {
	logger := Default()

	logger.name = name

	return logger
}

func Default() *Logger {
	logger := Logger{
		transports: []Appender{},
		level:      LogLevelAll,
	}

	return &logger
}

func (logger *Logger) SetLevel(level int) error {
	if level < LogLevelOff || level > LogLevelAll {
		return errors.New("invalid log level")
	}

	logger.level = level

	return nil
}
