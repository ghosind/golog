package logger

import "io"

type Appender interface {
	Output() io.Writer
}
