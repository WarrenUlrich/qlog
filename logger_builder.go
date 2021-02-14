package qlog

/*
	#include <windows.h>

	void setup_windows_console()
	{
		HANDLE hOut = GetStdHandle(STD_OUTPUT_HANDLE);
		DWORD dwMode = 0;
		GetConsoleMode(hOut, &dwMode);
		dwMode |= 0x0004;
		SetConsoleMode(hOut, dwMode);
	}
*/
import "C"

import (
	"io"
	"runtime"
	"sync"
)

//LoggerBuilder is a tool for building a logger.
//should be received through NewLoggerBuilder, and be built.
type LoggerBuilder struct {
	writers map[Level][]*struct {
		writer    io.Writer
		formatter Formatter
		mutex     sync.Mutex
	}

	async       bool
	workerCount int
}

//NewLoggerBuilder creates a new logger builder.
func NewLoggerBuilder() *LoggerBuilder {
	return &LoggerBuilder{}
}

//WithWriter adds a writer, as well as a formatter to write to it
//to the logger, it will only be written to on the specified events. (LevelInfo, LevelDebug, etc...)
func (l *LoggerBuilder) WithWriter(w io.Writer, f Formatter, levels ...Level) *LoggerBuilder {
	if l.writers == nil {
		l.writers = make(map[Level][]*struct {
			writer    io.Writer
			formatter Formatter
			mutex     sync.Mutex
		})
	}

	for _, level := range levels {
		temp := &struct {
			writer    io.Writer
			formatter Formatter
			mutex     sync.Mutex
		}{
			writer:    w,
			formatter: f,
		}

		l.writers[level] = append(l.writers[level], temp)
	}
	return l
}

//Async log messages will be handled in a seperate goroutine,
//workerCount specifies the amount of routines to create.
func (l *LoggerBuilder) Async(workerCount int) *LoggerBuilder {
	l.async = true
	l.workerCount = workerCount
	return l
}

//Build creates and returns a new logger.
func (l *LoggerBuilder) Build() Logger {

	//TODO: Move this to compile to call
	if runtime.GOOS == "windows" {
		C.setup_windows_console()
	}

	if l.async {
		return &asyncWriterLogger{
			writers:     l.writers,
			workerCount: l.workerCount,
		}
	}

	return &writerLogger{
		writers: l.writers,
	}
}

//BuildGlobal creates and returns a new logger.
//this logger will also be set as the global logger
//(accessed through qlog.Info, qlog.Debug, etc...).
func (l *LoggerBuilder) BuildGlobal() Logger {
	globalLogger = l.Build()
	return globalLogger
}

//BuildNamed creates and returns a new logger.
//this logger will also be added to the named loggers
//(accessed through qlog.Get("")). e.g. qlog.Get("console").Info("Hello world")
func (l *LoggerBuilder) BuildNamed(name string) Logger {
	if namedLoggers == nil {
		namedLoggers = make(map[string]Logger)
	}

	namedLoggers[name] = l.Build()
	return namedLoggers[name]
}
