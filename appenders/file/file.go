package file

import (
	"io/fs"
	"os"
	"sync"

	"github.com/ghosind/golog"
)

// Config is the configuration for the FileAppender.
type Config struct {
	// Filename is the log file name, it'll set to app.log if the value is empty.
	Filename string
	// Formatter is the log message formatter, default is TextFormatter.
	Formatter golog.Formatter
	// Logger is the logger instance.
	Logger *golog.Logger
	// Mode is the log file mode, default is 0644.
	Mode fs.FileMode
	// Path is the directory path to store logs, default is the current directory.
	Path string
}

// FileAppender is a golog appender that writes log into the specific file.
type FileAppender struct {
	file      *os.File
	filename  string
	formatter golog.Formatter
	logger    *golog.Logger
	mode      fs.FileMode
	mutex     sync.Mutex
	path      string
}

// New creates a new FileAppender.
func New(config ...Config) *FileAppender {
	cfg := Config{}
	if len(config) > 0 {
		cfg = config[0]
	}

	appender := FileAppender{
		logger: cfg.Logger,
	}

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

	if cfg.Formatter != nil {
		appender.formatter = cfg.Formatter
	} else {
		appender.formatter = &golog.TextFormatter{
			Logger: cfg.Logger,
		}
	}

	return &appender
}

// Write formats the data from entries and writes into the specific log file.
func (appender *FileAppender) Write(entry *golog.Entry) error {
	if appender.file == nil {
		if err := appender.openFile(); err != nil {
			return err
		}
	}

	buf, err := appender.formatter.Format(entry)
	if err != nil {
		return err
	}

	appender.file.Write(buf)

	return nil
}

// checkDir checks the directory path exists or not, if not exists, it will create the directory.
func (appender *FileAppender) checkDir() error {
	_, err := os.Stat(appender.path)
	if os.IsNotExist(err) {
		return os.Mkdir(appender.path, 0775)
	}

	return err
}

// openFile opens or creates the log file.
func (appender *FileAppender) openFile() error {
	appender.mutex.Lock()
	defer appender.mutex.Unlock()

	if err := appender.checkDir(); err != nil {
		return err
	}

	file, err := os.OpenFile(appender.filename, os.O_APPEND|os.O_CREATE|os.O_SYNC, appender.mode)
	if err != nil {
		return err
	}

	appender.file = file

	return nil
}
