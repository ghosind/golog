package stderr

import (
	"io"
	"os"
)

type StderrAppender struct {
}

func New() StderrAppender {
	return StderrAppender{}
}

func (appender *StderrAppender) Output() io.Writer {
	return os.Stderr
}
