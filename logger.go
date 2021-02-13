package qlog

//Logger ...
type Logger interface {
	//Info formats and writes a log message with LevelInfo severity.
	Info(args ...interface{}) error
	//Infof formats and writes a log message with LevelInfo severity.
	Infof(format string, args ...interface{}) error
	//Debug formats and writes a log message with LevelDebug severity.
	Debug(args ...interface{}) error
	//Debugf formats and writes a log message with LevelDebug severity.
	Debugf(format string, args ...interface{}) error
	//Warn formats and writes a log message with LevelWarn severity.
	Warn(args ...interface{}) error
	//Warnf formats and writes a log message with LevelWarn severity.
	Warnf(format string, args ...interface{}) error
	//Error formats and writes a log message with LevelError severity.
	Error(args ...interface{}) error
	//Errorf formats and writes a log message with LevelError severity.
	Errorf(format string, args ...interface{}) error
}
