package golog

import "fmt"

type TextFormatter struct {
	TimestampFormat string
}

func (formatter TextFormatter) Format(entry *Entry) ([]byte, error) {
	timestampFormat := formatter.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	time := entry.Time.Format(timestampFormat)

	return []byte(fmt.Sprintf("[%s] [%s] %s\n", time, GetLevel(entry.Level), entry.Message)), nil
}
