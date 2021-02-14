package qlog

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type writerLogger struct {
	writers map[Level][]*struct {
		writer    io.Writer
		formatter Formatter
		mutex     sync.Mutex
	}
}

func (w *writerLogger) write(level Level, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprint(args...),
	}

	for _, wf := range w.writers[level] {
		wf.mutex.Lock()
		wf.formatter.Format(
			&msg,
			wf.writer,
		)
		wf.mutex.Unlock()
	}

	return nil
}

func (w *writerLogger) writef(level Level, format string, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprintf(format, args...),
	}

	for _, wf := range w.writers[level] {
		wf.mutex.Lock()
		wf.formatter.Format(
			&msg,
			wf.writer,
		)
		wf.mutex.Unlock()
	}

	return nil
}

//Info writes a log message with info level severity.
func (w *writerLogger) Info(args ...interface{}) error {
	return w.write(LevelInfo, args...)
}

//Infof writes a log message with info level severity.
//Expects a string to format the provided args. e.g. qlog.Infof("%d", 49)
func (w *writerLogger) Infof(format string, args ...interface{}) error {
	return w.writef(LevelInfo, format, args...)
}

//Debug writes a log message with debug level severity.
func (w *writerLogger) Debug(args ...interface{}) error {
	return w.write(LevelDebug, args...)
}

//Debugf writes a log message with debug level severity.
//Expects a string to format the provided args. e.g. qlog.Debugf("%d", 49)
func (w *writerLogger) Debugf(format string, args ...interface{}) error {
	return w.writef(LevelDebug, format, args...)
}

//Warn writes a log message with warning level severity.
func (w *writerLogger) Warn(args ...interface{}) error {
	return w.write(LevelWarning, args...)
}

//Warnf writes a log message with warning level severity.
//Expects a string to format the provided args. e.g. qlog.Warnf("%d", 49)
func (w *writerLogger) Warnf(format string, args ...interface{}) error {
	return w.writef(LevelWarning, format, args...)
}

//Error writes a log message with error level severity.
func (w *writerLogger) Error(args ...interface{}) error {
	return w.write(LevelError, args...)
}

//Errorf writes a log message with error level severity.
//Expects a string to format the provided args. e.g. qlog.Errorf("%d", 49)
func (w *writerLogger) Errorf(format string, args ...interface{}) error {
	return w.writef(LevelError, format, args...)
}
