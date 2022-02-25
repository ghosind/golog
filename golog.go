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

func Panicln(args ...interface{}) {
	builtinLogger.Logln(PanicLevel, args...)
}

func Fatalln(args ...interface{}) {
	builtinLogger.Logln(FatalLevel, args...)
}

func Errorln(args ...interface{}) {
	builtinLogger.Logln(ErrorLevel, args...)
}

func Warnln(args ...interface{}) {
	builtinLogger.Logln(WarnLevel, args...)
}

func Infoln(args ...interface{}) {
	builtinLogger.Logln(InfoLevel, args...)
}

func Debugln(args ...interface{}) {
	builtinLogger.Logln(DebugLevel, args...)
}
