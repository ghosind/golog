package logger

func (logger *Logger) Fatal(fmt string, v ...interface{}) {
	if logger.level >= LogLevelFatal {
		logger.logger.Printf(fmt, v...)
	}
}

func (logger *Logger) Error(fmt string, v ...interface{}) {
	if logger.level >= LogLevelError {
		logger.logger.Printf(fmt, v...)
	}
}

func (logger *Logger) Warn(fmt string, v ...interface{}) {
	if logger.level >= LogLevelWarn {
		logger.logger.Printf(fmt, v...)
	}
}

func (logger *Logger) Info(fmt string, v ...interface{}) {
	if logger.level >= LogLevelInfo {
		logger.logger.Printf(fmt, v...)
	}
}

func (logger *Logger) Debug(fmt string, v ...interface{}) {
	if logger.level >= LogLevelDebug {
		logger.logger.Printf(fmt, v...)
	}
}

func (logger *Logger) Trace(fmt string, v ...interface{}) {
	if logger.level >= LogLevelTrace {
		logger.logger.Printf(fmt, v...)
	}
}
