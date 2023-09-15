package logger

type defaultLogger struct {
	// TODO: init
}

func (l *defaultLogger) InfoLog(v ...any) {
	// TODO: init
}

func (l *defaultLogger) ErrorLog(v ...any) {
	// TODO: init
}

func DefaultLogger() Logger {
	return &defaultLogger{}
}
