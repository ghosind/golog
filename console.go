package golog

import (
	"fmt"
	"os"
)

// ConsoleAppenderConfig is the configuration used to create a ConsoleAppender.
type ConsoleAppenderConfig struct {
	// Formatter is the formatter used to format the log entry.
	Formatter Formatter
}

// ConsoleAppender is a golog appender that writes log into the console.
type ConsoleAppender struct {
	formatter Formatter
}

// NewConsoleAppender creates a new ConsoleAppender.
func NewConsoleAppender(config ...ConsoleAppenderConfig) *ConsoleAppender {
	cfg := ConsoleAppenderConfig{}
	if len(config) > 0 {
		cfg = config[0]
	}

	appender := ConsoleAppender{}

	if cfg.Formatter != nil {
		appender.formatter = cfg.Formatter
	} else {
		appender.formatter = TextFormatter{}
	}

	return &appender
}

// Write formats the data from entries and writes into the console.
func (appender *ConsoleAppender) Write(entry *Entry) {
	fmt.Fprint(os.Stderr, string(appender.formatter.Format(entry)))
}
