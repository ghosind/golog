package console

import (
	"fmt"
	"os"

	"github.com/ghosind/golog"
)

// ConsoleAppenderConfig is the configuration used to create a ConsoleAppender.
type Config struct {
	// Formatter is the formatter used to format the log entry.
	Formatter golog.Formatter
}

// ConsoleAppender is a golog appender that writes log into the console.
type ConsoleAppender struct {
	formatter golog.Formatter
}

// New creates a new ConsoleAppender.
func New(config ...Config) *ConsoleAppender {
	cfg := Config{}
	if len(config) > 0 {
		cfg = config[0]
	}

	appender := ConsoleAppender{}

	if cfg.Formatter != nil {
		appender.formatter = cfg.Formatter
	} else {
		appender.formatter = golog.TextFormatter{}
	}

	return &appender
}

// Write formats the data from entries and writes into the console.
func (appender *ConsoleAppender) Write(entry *golog.Entry) {
	fmt.Fprint(os.Stderr, string(appender.formatter.Format(entry)))
}
