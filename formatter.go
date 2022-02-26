package golog

type Formatter interface {
	Format(entry *Entry) ([]byte, error)
}

const (
	defaultTimestampFormat string = "2006-01-02T15:04:05.000"
)

const (
	// KeyTimestamp is the timestamp field key.
	KeyTimestamp string = "timestamp"
	// KeyLevel is the level field key.
	KeyLevel string = "level"
	// KeyMessage is the message field key.
	KeyMessage string = "message"
)
