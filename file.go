package golog

import (
	"os"
	"sync"
)

type FileAppenderConfig struct {
	Path      string
	Filename  string
	Formatter Formatter
}

type FileAppender struct {
	path      string
	filename  string
	file      *os.File
	mutex     sync.Mutex
	formatter Formatter
}

func NewFileAppender(config FileAppenderConfig) *FileAppender {
	appender := FileAppender{}
	appender.path = config.Path
	appender.filename = config.Path + config.Filename
	appender.file = appender.openFile()

	if config.Formatter != nil {
		appender.formatter = config.Formatter
	} else {
		appender.formatter = TextFormatter{}
	}

	return &appender
}

func (appender *FileAppender) Write(entry *Entry) {
	if appender.file != nil {
		appender.file.Write(appender.formatter.Format(entry))
	}
}

func (appender *FileAppender) checkDir() error {
	if _, err := os.Stat(appender.path); os.IsNotExist(err) {
		return os.Mkdir(appender.path, 0777)
	}

	return nil
}

func (appender *FileAppender) openFile() *os.File {
	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	appender.checkDir()

	file, _ := os.OpenFile(appender.filename, os.O_APPEND|os.O_CREATE|os.O_SYNC, 0644)
	return file
}
