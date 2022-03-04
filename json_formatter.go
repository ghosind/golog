package golog

import "encoding/json"

type JSONFormatter struct {
	// Fields indicates custom fields' name.
	Fields map[string]string
	// TimestampFormat is the timestamp format template with time.Format rules.
	TimestampFormat string
	// Logger is the logger that the formatter is attached to.
	Logger *Logger
}

func (formatter JSONFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(map[string]interface{})

	timestampFormat := formatter.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	timestampField := KeyTimestamp
	if name := formatter.Fields[KeyTimestamp]; name != "" {
		timestampField = name
	}
	data[timestampField] = entry.Time.Format(timestampFormat)

	levelField := KeyLevel
	if name := formatter.Fields[KeyLevel]; name != "" {
		levelField = name
	}
	data[levelField] = GetLevelLabel(entry.Level)

	messageField := KeyMessage
	if name := formatter.Fields[KeyMessage]; name != "" {
		messageField = name
	}
	data[messageField] = entry.Message

	buf, err := json.Marshal(data)
	return buf, err
}
