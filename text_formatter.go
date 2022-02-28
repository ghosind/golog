package golog

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/ghosind/formatparser"
)

const (
	defaultTextFormat string = "[%timestamp] [%level] %message"
)

// TextFormatter is a formatter that format log messages to the specific format texts.
type TextFormatter struct {
	// LogFormat is the log message format template.
	LogFormat string
	// TimestampFormat is the timestamp format template with time.Format rules.
	TimestampFormat string

	formatParts []*formatparser.FormatPart
	mx          sync.Mutex
}

// Format formats message with the specific format.
func (formatter *TextFormatter) Format(entry *Entry) ([]byte, error) {
	parts := formatter.getFormatParts()
	buffer := bytes.Buffer{}

	timestampFormat := formatter.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	time := entry.Time.Format(timestampFormat)

	for _, part := range parts {
		if part.Type == formatparser.TypeText {
			buffer.WriteString(part.Value)
			continue
		} else if part.Type == formatparser.TypeUnknown {
			return nil, fmt.Errorf("unknown token %s", part.Value)
		}

		switch part.Value {
		case KeyTimestamp:
			buffer.WriteString(time)
		case KeyLevel:
			buffer.WriteString(GetLevel(entry.Level))
		case KeyMessage:
			buffer.WriteString(entry.Message)
		default:
			return nil, fmt.Errorf("unknown field %s", part.Value)
		}
	}

	return buffer.Bytes(), nil
}

// getFormatParts parses format template and returns template information.
func (formatter *TextFormatter) getFormatParts() []*formatparser.FormatPart {
	formatter.mx.Lock()
	defer formatter.mx.Unlock()

	if formatter.formatParts == nil {
		format := formatter.LogFormat
		if format == "" {
			format = defaultTextFormat
		}

		formatter.formatParts = formatparser.Parse(format)
	}

	return formatter.formatParts
}
