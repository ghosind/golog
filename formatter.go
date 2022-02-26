package golog

type Formatter interface {
	Format(entry *Entry) ([]byte, error)
}

const (
	defaultTimestampFormat string = "2006-01-02T15:04:05.000"
)
