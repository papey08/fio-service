package logger

import "fmt"

type Logger interface {
	InfoLog(s string)
	ErrorLog(s string)
	FatalLog(s string)
}

var l Logger

func InitLogger(logger Logger) {
	l = logger
}

func Info(format string, v ...any) {
	if len(v) == 0 {
		l.InfoLog(format)
	}
	l.InfoLog(fmt.Sprintf(format, v...))
}

func Error(format string, v ...any) {
	if len(v) == 0 {
		l.ErrorLog(format)
	}
	l.ErrorLog(fmt.Sprintf(format, v...))
}

func Fatal(format string, v ...any) {
	if len(v) == 0 {
		l.FatalLog(format)
	}
	l.FatalLog(fmt.Sprintf(format, v...))
}
