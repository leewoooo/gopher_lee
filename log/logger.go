package log

// Logger supports logging at various log levels.
type Logger interface {
	// Info
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infoln(v ...interface{})

	// warn
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnln(v ...interface{})

	// error
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})

	// fatal
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}
