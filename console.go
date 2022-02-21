package golog

import (
	"fmt"
	"os"
)

type ConsoleAppenderConfig struct {
	Formatter Formatter
}

type ConsoleAppender struct {
	Formatter Formatter
}

func NewConsoleAppender(config ConsoleAppenderConfig) *ConsoleAppender {
	appender := ConsoleAppender{}

	if config.Formatter != nil {
		appender.Formatter = config.Formatter
	} else {
		appender.Formatter = TextFormatter{}
	}

	return &appender
}

func (appender *ConsoleAppender) Write(entry *Entry) {
	fmt.Fprint(os.Stderr, string(appender.Formatter.Format(entry)))
}
