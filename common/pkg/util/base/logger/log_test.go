package logger

import "testing"

func TestSetLogger(t *testing.T) {
	l := &logger{level: LevelDebug}
	SetLogger(l)
}

func TestSetLevel(t *testing.T) {
	SetLevel(LevelAll)
	func() {
		defer func() {
			err := recover()
			if err != nil {
				t.Errorf("recorver returned err: %s", err)
			}
		}()
		SetLevel(1000)
	}()
}

func TestLogger_SetLevel(t *testing.T) {
	l := &logger{level: LevelDebug}
	l.SetLevel(LevelAll)
}

func TestLogger_Debug(t *testing.T) {
	l := &logger{level: LevelDebug}
	l.Debug("logger debug test")
}

func TestLogger_Info(t *testing.T) {
	l := &logger{level: LevelDebug}
	l.Info("logger info test")
}

func TestLogger_Warn(t *testing.T) {
	l := &logger{level: LevelDebug}
	l.Warn("logger warn test")
}

func TestLogger_Error(t *testing.T) {
	l := &logger{level: LevelDebug}
	l.Error("logger error test")
}

func TestDebug(t *testing.T) {
	Debug("log.Debug")
}

func TestInfo(t *testing.T) {
	Info("log.Info")
}

func TestWarn(t *testing.T) {
	Warn("log.Warn")
}

func TestError(t *testing.T) {
	Error("log.Error")
}
