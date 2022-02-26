package golog

import "encoding/json"

type JSONFormatter struct {
	TimestampFormat string
}

func (formatter JSONFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(map[string]interface{})

	timestampFormat := formatter.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	data["time"] = entry.Time.Format(timestampFormat)
	data["level"] = entry.Level
	data["message"] = entry.Message

	buf, err := json.Marshal(data)
	return buf, err
}
