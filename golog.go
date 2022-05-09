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

func Panicf(message string, args ...any) {
	builtinLogger.Panicf(message, args...)
}

func Fatalf(message string, args ...any) {
	builtinLogger.Fatalf(message, args...)
}

func Errorf(message string, args ...any) {
	builtinLogger.Errorf(message, args...)
}

func Warnf(message string, args ...any) {
	builtinLogger.Warnf(message, args...)
}

func Infof(message string, args ...any) {
	builtinLogger.Infof(message, args...)
}

func Debugf(message string, args ...any) {
	builtinLogger.Debugf(message, args...)
}

func Printf(message string, args ...any) {
	builtinLogger.Printf(message, args...)
}

func Panicln(args ...any) {
	builtinLogger.Panicln(args...)
}

func Fatalln(args ...any) {
	builtinLogger.Fatalln(args...)
}

func Errorln(args ...any) {
	builtinLogger.Errorln(args...)
}

func Warnln(args ...any) {
	builtinLogger.Warnln(args...)
}

func Infoln(args ...any) {
	builtinLogger.Infoln(args...)
}

func Debugln(args ...any) {
	builtinLogger.Debugln(args...)
}

func Println(args ...any) {
	builtinLogger.Println(args...)
}
