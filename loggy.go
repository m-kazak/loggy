package loggy

import (
	"log"
	"io"
	"fmt"
	"os"
)

// LogLevel type
type LogLevel int

// Log Level
const (
	FatalLevel LogLevel = iota + 1
	PanicLevel
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

// Info logs a message at Info level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Info(a ...interface{}) {
	l.internalLog(InfoLevel, false, fmt.Sprint(a...))
}

// Infof logs a message at Info level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Infof(format string, a ...interface{}) {
	l.internalLog(InfoLevel, false, fmt.Sprintf(format, a...))
}

// Infoln logs a message at Info level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Infoln(a ...interface{}) {
	l.internalLog(InfoLevel, true, fmt.Sprintln(a...))
}

// Warning logs a message at Warning level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Warning(a ...interface{}) {
	l.internalLog(WarningLevel, false, fmt.Sprint(a...))
}

// Warningf logs a message at Warning level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Warningf(format string, a ...interface{}) {
	l.internalLog(WarningLevel, false, fmt.Sprintf(format, a...))
}

// Warningln logs a message at Warning level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Warningln(a ...interface{}) {
	l.internalLog(WarningLevel, true, fmt.Sprintln(a...))
}

// Error logs a message at Error level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Error(a ...interface{}) {
	l.internalLog(ErrorLevel, false, fmt.Sprint(a...))
}

// Errorf logs a message at Error level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.internalLog(ErrorLevel, false, fmt.Sprintf(format, a...))
}

// Errorln logs a message at Error level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Errorln(a ...interface{}) {
	l.internalLog(ErrorLevel, true, fmt.Sprintln(a...))
}

// Panic logs a message at Panic level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Panic(a ...interface{}) {
	l.internalLog(PanicLevel, false, fmt.Sprint(a...))
	panic(fmt.Sprint(a...))
}

// Panicf logs a message at Panic level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Panicf(format string, a ...interface{}) {
	l.internalLog(PanicLevel, false, fmt.Sprintf(format, a...))
	panic(fmt.Sprintf(format, a...))
}

// Panicln logs a message at Panic level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Panicln(a ...interface{}) {
	l.internalLog(PanicLevel, true, fmt.Sprintln(a...))
	panic(fmt.Sprint(a...))
}

// Fatal logs a message at Fatal level. Arguments are handled in the manner of fmt.Print
func (l *Logger) Fatal(a ...interface{}) {
	l.internalLog(FatalLevel, false, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatal logs a message at Fatal level. Arguments are handled in the manner of fmt.Printf
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.internalLog(FatalLevel, false, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalln logs a message at Fatal level. Arguments are handled in the manner of fmt.Println
func (l *Logger) Fatalln(a ...interface{}) {
	l.internalLog(FatalLevel, true, fmt.Sprintln(a...))
	os.Exit(1)
}

// Return LogLeve as a string
func getLevelString (level LogLevel) string {
	logLevels := [...]string{
		"[FATAL]",
		"[PANIC]",
		"[ERROR]",
		"[WARNING]",
		"[INFO]",
		"[DEBUG]",
	}
	return logLevels[level-1]
}