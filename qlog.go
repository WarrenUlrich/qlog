package qlog

var (
	globalLogger Logger
	namedLoggers map[string]Logger
)

//Info ...
func Info(args ...interface{}) error {
	return globalLogger.Info(args...)
}

//Infof ...
func Infof(format string, args ...interface{}) error {
	return globalLogger.Infof(format, args...)
}

//Debug ...
func Debug(args ...interface{}) error {
	return globalLogger.Debug(args...)
}

//Debugf ...
func Debugf(format string, args ...interface{}) error {
	return globalLogger.Debugf(format, args...)
}

//Warn ...
func Warn(args ...interface{}) error {
	return globalLogger.Warn(args...)
}

//Warnf ...
func Warnf(format string, args ...interface{}) error {
	return globalLogger.Warnf(format, args...)
}

//Error ...
func Error(args ...interface{}) error {
	return globalLogger.Error(args...)
}

//Errorf ...
func Errorf(format string, args ...interface{}) error {
	return globalLogger.Errorf(format, args...)
}

//Get ...
func Get(name string) Logger {
	l, ok := namedLoggers[name]
	if ok {
		return l
	}

	return nil
}