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

// GetLevel returns the string representation of the level.
func GetLevel(level Level) string {
	switch level {
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	case ErrorLevel:
		return "ERROR"
	case WarnLevel:
		return "WARN"
	case InfoLevel:
		return "INFO"
	case DebugLevel:
		return "DEBUG"
	default:
		return ""
	}
}
