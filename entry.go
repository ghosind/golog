package golog

import (
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Level   Level
	Message string
	Time    time.Time

	logger *Logger
}

func newEntry() *Entry {
	return &Entry{}
}

func (entry *Entry) Log(level Level, message string) {
	entry.Level = level
	entry.Message = message
	entry.Time = time.Now()

	for _, appender := range entry.logger.Appenders {
		appender.Write(entry)
	}

	if level == FatalLevel {
		// exit program if level is FatalLevel
		os.Exit(1)
	} else if level == PanicLevel {
		// trigger panic if level is PanicLevel
		panic(message)
	}
}

func (entry *Entry) Logf(level Level, message string, args ...interface{}) {
	entry.Log(level, fmt.Sprintf(message, args...))
}

func (entry *Entry) Logln(level Level, args ...interface{}) {
	msg := fmt.Sprintln(args...)
	entry.Log(level, msg[:len(msg)-1])
}

func (entry *Entry) Panic(message string) {
	entry.Log(PanicLevel, message)
}

func (entry *Entry) Fatal(message string) {
	entry.Log(FatalLevel, message)
}

func (entry *Entry) Error(message string) {
	entry.Log(ErrorLevel, message)
}

func (entry *Entry) Warn(message string) {
	entry.Log(WarnLevel, message)
}

func (entry *Entry) Info(message string) {
	entry.Log(InfoLevel, message)
}

func (entry *Entry) Debug(message string) {
	entry.Log(DebugLevel, message)
}

func (entry *Entry) Panicf(message string, args ...interface{}) {
	entry.Logf(PanicLevel, message, args...)
}

func (entry *Entry) Fatalf(message string, args ...interface{}) {
	entry.Logf(FatalLevel, message, args...)
}

func (entry *Entry) Errorf(message string, args ...interface{}) {
	entry.Logf(ErrorLevel, message, args...)
}

func (entry *Entry) Warnf(message string, args ...interface{}) {
	entry.Logf(WarnLevel, message, args...)
}

func (entry *Entry) Infof(message string, args ...interface{}) {
	entry.Logf(InfoLevel, message, args...)
}

func (entry *Entry) Debugf(message string, args ...interface{}) {
	entry.Logf(DebugLevel, message, args...)
}

func (entry *Entry) Panicln(args ...interface{}) {
	entry.Logln(PanicLevel, args...)
}

func (entry *Entry) Fatalln(args ...interface{}) {
	entry.Logln(FatalLevel, args...)
}

func (entry *Entry) Errorln(args ...interface{}) {
	entry.Logln(ErrorLevel, args...)
}

func (entry *Entry) Warnln(args ...interface{}) {
	entry.Logln(WarnLevel, args...)
}

func (entry *Entry) Infoln(args ...interface{}) {
	entry.Logln(InfoLevel, args...)
}

func (entry *Entry) Debugln(args ...interface{}) {
	entry.Logln(DebugLevel, args...)
}
