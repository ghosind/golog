package file

import (
	"io"
	"os"
)

type FileAppender struct {
	output io.Writer

	path string
}

func New(path string) (*FileAppender, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	return &FileAppender{
		output: file,
		path:   path,
	}, nil
}

func (appender *FileAppender) Output() io.Writer {
	return appender.output
}
