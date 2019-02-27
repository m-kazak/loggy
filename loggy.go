package loggy

import (
	"log"
	"io"
	"fmt"
)

// LogLevel type
type LogLevel int

// Log Level
const (
	CriticalLevel LogLevel = iota + 1
	ErrorLevel
	WarningLevel
	InfoLevel
	DebugLevel
)

type Logger struct {
	// Original GoLang logger
	worker *log.Logger

	// Current level of logging
	level LogLevel
}

// Returns a new instance of Logger class
func New(out io.Writer, flag int, initLevel LogLevel) *Logger {
	return  &Logger{worker: log.New(out, "", flag), level: initLevel}
}

// internalLog actually making log
func (l *Logger) internalLog (level LogLevel, ln bool, input string) {
	if level > l.level {
		return
	}

	msg := fmt.Sprintf("%s %s", getLevelString(level), input)

	if ln {
		l.worker.Print(msg)
	} else {
		l.worker.Println(msg)
	}
	
}

// Debug logs a message at Debug level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Debug(a ...interface{}) {
	l.internalLog(DebugLevel, false, fmt.Sprint(a...))
}

// Debugf logs a message at Debug level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.internalLog(DebugLevel, false, fmt.Sprintf(format, a...))
}

// Debugln logs a message at Debug level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Debugln(a ...interface{}) {
	l.internalLog(DebugLevel, true, fmt.Sprintln(a...))
}

// Return LogLeve as a string
func getLevelString (level LogLevel) string {
	logLevels := [...]string{
		"[CRITICAL]",
		"[ERROR]",
		"[WARNING]",
		"[INFO]",
		"[DEBUG]",
	}
	return logLevels[level-1]
}