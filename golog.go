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

func Print(message string) {
	builtinLogger.Print(message)
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

func Printf(message string, args ...interface{}) {
	builtinLogger.Printf(message, args...)
}

func Panicln(args ...interface{}) {
	builtinLogger.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	builtinLogger.Fatalln(args...)
}

func Errorln(args ...interface{}) {
	builtinLogger.Errorln(args...)
}

func Warnln(args ...interface{}) {
	builtinLogger.Warnln(args...)
}

func Infoln(args ...interface{}) {
	builtinLogger.Infoln(args...)
}

func Debugln(args ...interface{}) {
	builtinLogger.Debugln(args...)
}

func Println(args ...interface{}) {
	builtinLogger.Println(args...)
}
