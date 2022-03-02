package golog

import "testing"

func Test_GetLevelLabel(t *testing.T) {
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
		if actual := GetLevelLabel(level); actual != message {
			t.Errorf("GetLevelLabel(%d) = %s, expect %s", level, actual, message)
		}
	}
}
