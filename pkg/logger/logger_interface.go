package logger

type Logger interface {
	InfoLog(v ...any)
	ErrorLog(v ...any)
}

var l Logger

func InitLogger(logger Logger) {
	l = logger
}

func Info(v ...any) {
	l.InfoLog(v)
}

func Error(v ...any) {
	l.ErrorLog()
}
