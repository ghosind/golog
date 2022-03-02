package golog

import (
	"fmt"
	"testing"
	"time"
)

func Test_TextFormatterWithDefaultFormat(t *testing.T) {
	entry := Entry{
		Level:   InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}
	formatter := TextFormatter{}

	buf, err := formatter.Format(&entry)
	if err != nil {
		t.Errorf("TextFormatter Format() returns %v error, expect no error", err)
	}

	expect := fmt.Sprintf("[%s] [%s] %s",
		entry.Time.Format(defaultTimestampFormat),
		GetLevelLabel(entry.Level),
		entry.Message,
	)
	if string(buf) != expect {
		t.Errorf("TextFormatter Format() returns \"%s\", expect \"%s\"", string(buf), expect)
	}
}

func Test_TextFormatterWithCustomFormat(t *testing.T) {
	entry := Entry{
		Level:   InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}
	timestampFormat := "2006-01-02 15:04:05"
	formatter := TextFormatter{
		LogFormat:       "%level - %timestamp - %message",
		TimestampFormat: timestampFormat,
	}

	buf, err := formatter.Format(&entry)
	if err != nil {
		t.Errorf("TextFormatter Format() returns %v error, expect no error", err)
	}

	expect := fmt.Sprintf("%s - %s - %s",
		GetLevelLabel(entry.Level),
		entry.Time.Format(timestampFormat),
		entry.Message,
	)
	if string(buf) != expect {
		t.Errorf("TextFormatter Format() returns \"%s\", expect \"%s\"", string(buf), expect)
	}
}
