package qlog

var (
	globalLogger Logger
	namedLoggers map[string]Logger
)

//Info formats and writes a log message with LevelInfo severity.
func Info(args ...interface{}) error {
	return globalLogger.Info(args...)
}

//Infof formats and writes a log message with LevelInfo severity.
func Infof(format string, args ...interface{}) error {
	return globalLogger.Infof(format, args...)
}

//Debug formats and writes a log message with LevelDebug severity.
func Debug(args ...interface{}) error {
	return globalLogger.Debug(args...)
}

//Debugf formats and writes a log message with LevelDebug severity.
func Debugf(format string, args ...interface{}) error {
	return globalLogger.Debugf(format, args...)
}

//Warn formats and writes a log message with LevelWarn severity.
func Warn(args ...interface{}) error {
	return globalLogger.Warn(args...)
}

//Warnf formats and writes a log message with LevelWarn severity.
func Warnf(format string, args ...interface{}) error {
	return globalLogger.Warnf(format, args...)
}

//Error formats and writes a log message with LevelError severity.
func Error(args ...interface{}) error {
	return globalLogger.Error(args...)
}

//Errorf formats and writes a log message with LevelError severity.
func Errorf(format string, args ...interface{}) error {
	return globalLogger.Errorf(format, args...)
}

//Get obtains a named logger, returns nil if it isn't found.
func Get(name string) Logger {
	l, ok := namedLoggers[name]
	if ok {
		return l
	}

	return nil
}