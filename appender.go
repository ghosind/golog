package golog

import (
	"fmt"
	"os"
)

type Appender interface {
	Write(entry *Entry) error
}

// basicAppender is a basic appender that writes log into the console. It is used for default logger.
type basicAppender struct {
	formatter Formatter
}

// newBasicAppender creates a new basic appender.
func newBasicAppender() *basicAppender {
	return &basicAppender{
		formatter: &TextFormatter{},
	}
}

// Write formats the data from entries by text formatter and writes into the console.
func (appender *basicAppender) Write(entry *Entry) error {
	buf, err := appender.formatter.Format(entry)
	if err != nil {
		return err
	}

	fmt.Fprint(os.Stderr, string(buf))
	return nil
}
