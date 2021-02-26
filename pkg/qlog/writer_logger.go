package qlog

import (
	"fmt"
	"io"
	"time"
)

//Logger ...
type Logger struct {
	writers map[Level][]*struct {
		writer    io.Writer
		formatter Formatter
	}
}

func (w *Logger) write(level Level, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprint(args...),
	}

	for _, wf := range w.writers[level] {
		wf.formatter.Format(
			&msg,
			wf.writer,
		)
	}

	return nil
}

func (w *Logger) writef(level Level, format string, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprintf(format, args...),
	}

	for _, wf := range w.writers[level] {
		wf.formatter.Format(
			&msg,
			wf.writer,
		)
	}

	return nil
}

//Info writes a log message with info level severity.
func (w *Logger) Info(args ...interface{}) error {
	return w.write(LevelInfo, args...)
}

//Infof writes a log message with info level severity.
//Expects a string to format the provided args. e.g. qlog.Infof("%d", 49)
func (w *Logger) Infof(format string, args ...interface{}) error {
	return w.writef(LevelInfo, format, args...)
}

//Debug writes a log message with debug level severity.
func (w *Logger) Debug(args ...interface{}) error {
	return w.write(LevelDebug, args...)
}

//Debugf writes a log message with debug level severity.
//Expects a string to format the provided args. e.g. qlog.Debugf("%d", 49)
func (w *Logger) Debugf(format string, args ...interface{}) error {
	return w.writef(LevelDebug, format, args...)
}

//Warn writes a log message with warning level severity.
func (w *Logger) Warn(args ...interface{}) error {
	return w.write(LevelWarning, args...)
}

//Warnf writes a log message with warning level severity.
//Expects a string to format the provided args. e.g. qlog.Warnf("%d", 49)
func (w *Logger) Warnf(format string, args ...interface{}) error {
	return w.writef(LevelWarning, format, args...)
}

//Error writes a log message with error level severity.
func (w *Logger) Error(args ...interface{}) error {
	return w.write(LevelError, args...)
}

//Errorf writes a log message with error level severity.
//Expects a string to format the provided args. e.g. qlog.Errorf("%d", 49)
func (w *Logger) Errorf(format string, args ...interface{}) error {
	return w.writef(LevelError, format, args...)
}
