package golog

import (
	"io/fs"
	"os"
	"sync"
)

// FileAppenderConfig is the configuration for the FileAppender.
type FileAppenderConfig struct {
	// Path is the directory path to store logs, default is the current directory.
	Path string
	// Filename is the log file name, it'll set to app.log if the value is empty.
	Filename string
	// Mode is the log file mode, default is 0644.
	Mode fs.FileMode
	// Formatter is the log message formatter, default is TextFormatter.
	Formatter Formatter
}

// FileAppender is a golog appender that writes log into the specific file.
type FileAppender struct {
	path      string
	filename  string
	file      *os.File
	mode      fs.FileMode
	mutex     sync.Mutex
	formatter Formatter
}

// NewFileAppender creates a new FileAppender.
func NewFileAppender(config ...FileAppenderConfig) *FileAppender {
	cfg := FileAppenderConfig{}
	if len(config) > 0 {
		cfg = config[0]
	}

	appender := FileAppender{}

	appender.path = cfg.Path
	appender.filename = cfg.Filename
	if appender.path == "" {
		appender.path = "."
	}
	if appender.filename == "" {
		appender.filename = "app.log"
	}
	appender.filename = appender.path + "/" + appender.filename

	if cfg.Mode == 0 {
		appender.mode = 0644
	} else {
		appender.mode = cfg.Mode
	}

	appender.file = appender.openFile()

	if cfg.Formatter != nil {
		appender.formatter = cfg.Formatter
	} else {
		appender.formatter = TextFormatter{}
	}

	return &appender
}

// Write formats the data from entries and writes into the specific log file.
func (appender *FileAppender) Write(entry *Entry) {
	if appender.file != nil {
		appender.file.Write(appender.formatter.Format(entry))
	}
}

// checkDir checks the directory path exists or not, if not exists, it will create the directory.
func (appender *FileAppender) checkDir() error {
	_, err := os.Stat(appender.path)
	if os.IsNotExist(err) {
		return os.Mkdir(appender.path, 0775)
	} else if err != nil {
		panic(err)
	}

	return nil
}

// openFile opens or creates the log file.
func (appender *FileAppender) openFile() *os.File {
	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	appender.checkDir()

	file, err := os.OpenFile(appender.filename, os.O_APPEND|os.O_CREATE|os.O_SYNC, appender.mode)
	if err != nil {
		panic(err)
	}

	return file
}
