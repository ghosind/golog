package stdout

import (
	"io"
	"os"
)

type StdoutAppender struct {
}

func New() StdoutAppender {
	return StdoutAppender{}
}

func (appender *StdoutAppender) Output() io.Writer {
	return os.Stdout
}
