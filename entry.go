package golog

import "time"

type Entry struct {
	logger  *Logger
	level   Level
	message string
	time    time.Time
}

func newEntry() *Entry {
	return &Entry{}
}

func (entry *Entry) Log(level Level, message string) {
	for _, appender := range entry.logger.Appenders {
		appender.Write(entry)
	}
}
