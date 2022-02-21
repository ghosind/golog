package golog

import (
	"encoding/json"
	"fmt"
)

type Formatter interface {
	Format(entry *Entry) []byte
}

type TextFormatter struct{}

func (formatter TextFormatter) Format(entry *Entry) []byte {
	time := entry.time.Format("2006-01-02T15:04:05.000")

	return []byte(fmt.Sprintf("[%s] [%s] %s", time, GetLevel(entry.level), entry.message))
}

type JSONFormatter struct{}

func (formatter JSONFormatter) Format(entry *Entry) []byte {
	data := make(map[string]interface{})
	data["time"] = entry.time.Format("2006-01-02T15:04:05.000")
	data["level"] = entry.level
	data["message"] = entry.message

	re, _ := json.Marshal(data)
	return re
}
