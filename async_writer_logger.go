package qlog

import (
	"fmt"
	"io"
	"runtime"
	"sync"
	"time"
)

type asyncWriterLogger struct {
	writers map[Level][]*struct {
		writer    io.Writer
		formatter Formatter
		mutex     sync.Mutex
	}

	workerCount int
	channel     chan *Message
}

func (a *asyncWriterLogger) startup() error {
	if a.channel == nil {
		a.channel = make(chan *Message)
		fmt.Println(a.workerCount)
		for i := 0; i < a.workerCount; i++ {
			go func() {
				for {
					if msg, ok := <-a.channel; ok {
						for _, wf := range a.writers[msg.Level] {
							wf.mutex.Lock()
							wf.formatter.Format(
								msg,
								wf.writer,
							)
							wf.mutex.Unlock()
						}
					} else {
						break
					}
				}
			}()
		}

		runtime.SetFinalizer(a, func(a *asyncWriterLogger) {
			close(a.channel)
		})
	}
	return nil
}

func (a *asyncWriterLogger) write(level Level, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprint(args...),
	}

	a.startup()
	a.channel <- &msg
	return nil
}

func (a *asyncWriterLogger) writef(level Level, format string, args ...interface{}) error {
	msg := Message{
		Level:   level,
		Time:    time.Now(),
		Message: fmt.Sprintf(format, args...),
	}

	a.startup()
	a.channel <- &msg
	return nil
}

//Info writes a log message with info level severity.
func (a *asyncWriterLogger) Info(args ...interface{}) error {
	return a.write(LevelInfo, args...)
}

//Infof writes a log message with info level severity.
//Expects a string to format the provided args. e.g. qlog.Infof("%d", 49)
func (a *asyncWriterLogger) Infof(format string, args ...interface{}) error {
	return a.writef(LevelInfo, format, args...)
}

//Debug writes a log message with debug level severity.
func (a *asyncWriterLogger) Debug(args ...interface{}) error {
	return a.write(LevelDebug, args...)
}

//Debugf writes a log message with debug level severity.
//Expects a string to format the provided args. e.g. qlog.Debugf("%d", 49)
func (a *asyncWriterLogger) Debugf(format string, args ...interface{}) error {
	return a.writef(LevelDebug, format, args...)
}

//Warn writes a log message with warning level severity.
func (a *asyncWriterLogger) Warn(args ...interface{}) error {
	return a.write(LevelWarning, args...)
}

//Warnf writes a log message with warning level severity.
//Expects a string to format the provided args. e.g. qlog.Warnf("%d", 49)
func (a *asyncWriterLogger) Warnf(format string, args ...interface{}) error {
	return a.writef(LevelWarning, format, args...)
}

//Error writes a log message with error level severity.
func (a *asyncWriterLogger) Error(args ...interface{}) error {
	return a.write(LevelError, args...)
}

//Errorf writes a log message with error level severity.
//Expects a string to format the provided args. e.g. qlog.Errorf("%d", 49)
func (a *asyncWriterLogger) Errorf(format string, args ...interface{}) error {
	return a.writef(LevelError, format, args...)
}
