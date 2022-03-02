package golog

// Level represents the log level.
type Level int

const (
	// PanicLevel level, the highest level of severity. It'll logs and then calls panic with the message.
	PanicLevel Level = iota
	// FatalLevel level, critical level of severity. It'll logs and then calls os.Exit(1).
	FatalLevel
	// ErrorLevel level, error conditions.
	ErrorLevel
	// WarnLevel level, warning but not an error.
	WarnLevel
	// InfoLevel level, informational messages.
	InfoLevel
	// DebugLevel level, debug-level messages are typically voluminous to debug an application.
	DebugLevel
)

var defaultLevelLabels map[Level]string = map[Level]string{
	PanicLevel: "PANIC",
	FatalLevel: "FATAL",
	ErrorLevel: "ERROR",
	WarnLevel:  "WARN",
	InfoLevel:  "INFO",
	DebugLevel: "DEBUG",
}

// GetLevelLabel returns the string representation of the level.
func GetLevelLabel(level Level) string {
	return builtinLogger.GetLevelLabel(level)
}
