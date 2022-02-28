package console

import (
	"fmt"
	"io"
	"os"

	"github.com/ghosind/golog"
)

// ConsoleAppenderConfig is the configuration used to create a ConsoleAppender.
type Config struct {
	// Formatter is the formatter used to format the log entry.
	Formatter golog.Formatter
	// Output is the logger output file descriptor, default is Stderr.
	Output io.Writer
}

// ConsoleAppender is a golog appender that writes log into the console.
type ConsoleAppender struct {
	formatter golog.Formatter
	output    io.Writer
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
		appender.formatter = &golog.TextFormatter{}
	}

	if cfg.Output != nil {
		appender.output = cfg.Output
	} else {
		appender.output = os.Stderr
	}

	return &appender
}

// Write formats the data from entries and writes into the console.
func (appender *ConsoleAppender) Write(entry *golog.Entry) error {
	buf, err := appender.formatter.Format(entry)
	if err != nil {
		return err
	}
	fmt.Fprint(appender.output, string(buf))

	return nil
}
