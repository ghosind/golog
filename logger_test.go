package golog

import (
	"testing"
)

func Test_LoggerGetLevelLabel(t *testing.T) {
	expects := map[Level]string{
		PanicLevel: "PANIC",
		FatalLevel: "FATAL",
		ErrorLevel: "ERROR",
		WarnLevel:  "WARN",
		InfoLevel:  "INFO",
		DebugLevel: "DEBUG",
		-1:         "",
	}
	logger := New()

	for level, message := range expects {
		if actual := logger.GetLevelLabel(level); actual != message {
			t.Errorf("GetLevelLabel(%d) = %s, expect %s", level, actual, message)
		}
	}

	expects = map[Level]string{
		PanicLevel: "Panic",
		FatalLevel: "Fatal",
		ErrorLevel: "Error",
		WarnLevel:  "Warning",
		InfoLevel:  "info",
		DebugLevel: "debug",
		-1:         "",
	}
	logger.LevelLabels = expects

	for level, message := range expects {
		if actual := logger.GetLevelLabel(level); actual != message {
			t.Errorf("GetLevelLabel(%d) = %s, expect %s", level, actual, message)
		}
	}
}

func assertLevelEnable(t *testing.T, expect bool, actual bool) {
	if expect != actual {
		t.Errorf("Expect %v, actual %v", expect, actual)
	}
}

func Test_LoggerLevelEnable(t *testing.T) {
	logger := New()

	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = PanicLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = FatalLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = ErrorLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = WarnLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = InfoLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, false, logger.isLevelEnabled(DebugLevel))

	logger.Level = DebugLevel
	assertLevelEnable(t, true, logger.isLevelEnabled(PanicLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(FatalLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(ErrorLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(WarnLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(InfoLevel))
	assertLevelEnable(t, true, logger.isLevelEnabled(DebugLevel))
}
