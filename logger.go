package smtp

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	// Silent silent log level
	Silent LogLevel = iota + 1
	// Error error log level
	Error
	// Warn warn log level
	Warn
	// Info info log level
	Info
	// Debug debug log level
	Debug
)

// Logger interface is used by Server to report unexpected internal errors and other debug information.
type Logger interface {
	LogMode(level LogLevel) Logger
	Error(err error, msg string, v ...interface{})
	Info(msg string, v ...interface{})
	Debug(msg string, v ...interface{})
	Warn(msg string, v ...interface{})
}

type logger struct {
	l     *log.Logger
	level LogLevel
}

func (l logger) LogMode(level LogLevel) Logger {
	l.level = level
	return l
}

func (l logger) Error(err error, msg string, v ...interface{}) {
	if l.level < Error {
		return
	}
	l.l.Printf(fmt.Sprintf("error - %s", msg), v...)
	l.l.Println(err)
}

func (l logger) Info(msg string, v ...interface{}) {
	if l.level < Info {
		return
	}
	l.l.Printf(fmt.Sprintf("info - %s", msg), v...)
}

func (l logger) Debug(msg string, v ...interface{}) {
	if l.level < Debug {
		return
	}
	l.l.Printf(fmt.Sprintf("debug - %s", msg), v...)
}

func (l logger) Warn(msg string, v ...interface{}) {
	if l.level < Warn {
		return
	}
	l.l.Printf(fmt.Sprintf("warn - %s", msg), v...)
}

func createLogger() Logger {
	return &logger{
		l:     log.New(os.Stdout, "smtp/server ", log.LstdFlags),
		level: Info,
	}
}
