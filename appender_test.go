package golog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func Test_BasicAppender(t *testing.T) {
	stderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	entry := Entry{
		Level:   InfoLevel,
		Time:    time.Now(),
		Message: "Hello World",
	}
	formatter := newBasicAppender()

	if err := formatter.Write(&entry); err != nil {
		t.Errorf("basicFormatter.Write returns %v error, expect no error", err)
	}

	w.Close()
	os.Stderr = stderr

	buf := bytes.Buffer{}
	io.Copy(&buf, r)

	expect := fmt.Sprintf("[%s] [%s] %s",
		entry.Time.Format(defaultTimestampFormat),
		GetLevel(entry.Level),
		entry.Message,
	)

	if buf.String() != expect {
		t.Errorf("basicFormatter logs \"%s\", expect \"%s\"", buf.String(), expect)
	}
}
