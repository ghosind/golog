package golog

type Appender interface {
	Write(entry *Entry)
}
