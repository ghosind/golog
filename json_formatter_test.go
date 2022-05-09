package golog

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

func Test_DefaultJSONFormatter(t *testing.T) {
	entry := Entry{
		Level:   InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}

	expect, err := json.Marshal(map[string]any{
		KeyTimestamp: entry.Time.Format(defaultTimestampFormat),
		KeyLevel:     GetLevelLabel(entry.Level),
		KeyMessage:   entry.Message,
	})
	if err != nil {
		log.Fatalf("Failed to marshal expect json: %v", err)
	}

	formatter := JSONFormatter{}

	buf, err := formatter.Format(&entry)
	if err != nil {
		t.Errorf("JSONFormatter Format() returns %v error, expect no error", err)
	}

	if string(expect) != string(buf) {
		t.Errorf("TextFormatter Format() returns \"%s\", expect \"%s\"", string(buf), string(expect))
	}
}

func Test_CustomJSONFormatter(t *testing.T) {
	entry := Entry{
		Level:   InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}

	timestampFormat := "2006-01-02 15:04:05"
	expect, err := json.Marshal(map[string]any{
		"TIME":    entry.Time.Format(timestampFormat),
		"LEVEL":   GetLevelLabel(entry.Level),
		"MESSAGE": entry.Message,
	})
	if err != nil {
		log.Fatalf("Failed to marshal expect json: %v", err)
	}

	formatter := JSONFormatter{
		TimestampFormat: timestampFormat,
		Fields: map[string]string{
			KeyLevel:     "LEVEL",
			KeyTimestamp: "TIME",
			KeyMessage:   "MESSAGE",
		},
	}

	buf, err := formatter.Format(&entry)
	if err != nil {
		t.Errorf("JSONFormatter Format() returns %v error, expect no error", err)
	}

	if string(expect) != string(buf) {
		t.Errorf("TextFormatter Format() returns \"%s\", expect \"%s\"", string(buf), string(expect))
	}
}
