package golog

var builtinLogger *Logger = Default()

func Panic(message string) {
	builtinLogger.Panic(message)
}

func Fatal(message string) {
	builtinLogger.Fatal(message)
}

func Error(message string) {
	builtinLogger.Error(message)
}

func Warn(message string) {
	builtinLogger.Warn(message)
}

func Info(message string) {
	builtinLogger.Info(message)
}

func Debug(message string) {
	builtinLogger.Debug(message)
}

func Panicf(message string, args ...interface{}) {
	builtinLogger.Panicf(message, args...)
}

func Fatalf(message string, args ...interface{}) {
	builtinLogger.Fatalf(message, args...)
}

func Errorf(message string, args ...interface{}) {
	builtinLogger.Errorf(message, args...)
}

func Warnf(message string, args ...interface{}) {
	builtinLogger.Warnf(message, args...)
}

func Infof(message string, args ...interface{}) {
	builtinLogger.Infof(message, args...)
}

func Debugf(message string, args ...interface{}) {
	builtinLogger.Debugf(message, args...)
}
