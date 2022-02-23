package golog

import "testing"

func Test_GetLevel(t *testing.T) {
	expects := map[Level]string{
		PanicLevel: "PANIC",
		FatalLevel: "FATAL",
		ErrorLevel: "ERROR",
		WarnLevel:  "WARN",
		InfoLevel:  "INFO",
		DebugLevel: "DEBUG",
		-1:         "",
	}

	for level, message := range expects {
		if actual := GetLevel(level); actual != message {
			t.Errorf("GetLevel(%d) = %s, expect %s", level, actual, message)
		}
	}
}
