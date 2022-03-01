package console

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/ghosind/golog"
)

func Test_StdConsoleAppender(t *testing.T) {
	stderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	entry := golog.Entry{
		Level:   golog.InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}
	appender := New()

	if err := appender.Write(&entry); err != nil {
		t.Errorf("ConsoleAppender.Write returns %v error, expect no error", err)
	}

	w.Close()
	os.Stderr = stderr

	buf := bytes.Buffer{}
	io.Copy(&buf, r)

	expect := fmt.Sprintf("[%s] [%s] %s",
		entry.Time.Format("2006-01-02T15:04:05.000"),
		golog.GetLevel(entry.Level),
		entry.Message,
	)

	if buf.String() != expect {
		t.Errorf("ConsoleAppender logs \"%s\", expect \"%s\"", buf.String(), expect)
	}
}

func Test_CustomConsoleAppender(t *testing.T) {
	buf := bytes.Buffer{}

	appender := New(Config{
		Formatter: &golog.TextFormatter{
			LogFormat:       "[%level] [%timestamp] %message",
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Output: &buf,
	})

	entry := golog.Entry{
		Level:   golog.InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}

	if err := appender.Write(&entry); err != nil {
		t.Errorf("Custom ConsoleAppender.Write returns %v error, expect no error", err)
	}

	expect := fmt.Sprintf("[%s] [%s] %s",
		golog.GetLevel(entry.Level),
		entry.Time.Format("2006-01-02 15:04:05"),
		entry.Message,
	)

	if buf.String() != expect {
		t.Errorf("Custom ConsoleAppender logs \"%s\", expect \"%s\"", buf.String(), expect)
	}
}
