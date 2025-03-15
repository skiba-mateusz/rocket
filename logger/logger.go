package logger

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

type Logger interface {
	Debug(msg string, args ...interface{})
	Success(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type LogLevel int

const (
	DEBUG = iota
	INFO
	SUCCESS
	WARN
	ERROR
)

type DefaultLogger struct {
	level      LogLevel
	logger     *log.Logger
	colorFuncs map[LogLevel]func(a ...interface{}) string
}

func NewDefaultLogger(level LogLevel) *DefaultLogger {
	colorFuncs := map[LogLevel]func(a ...interface{}) string{
		DEBUG:   color.New(color.FgMagenta).SprintFunc(),
		SUCCESS: color.New(color.FgGreen).SprintFunc(),
		INFO:    color.New(color.FgBlue).SprintFunc(),
		WARN:    color.New(color.FgYellow).SprintFunc(),
		ERROR:   color.New(color.FgRed).SprintFunc(),
	}

	return &DefaultLogger{
		level:      level,
		logger:     log.New(os.Stdout, "", 0),
		colorFuncs: colorFuncs,
	}
}

func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
	l.logMessage(DEBUG, "DEBUG", msg, args...)
}

func (l *DefaultLogger) Info(msg string, args ...interface{}) {
	l.logMessage(INFO, "INFO", msg, args...)
}

func (l *DefaultLogger) Success(msg string, args ...interface{}) {
	l.logMessage(SUCCESS, "SUCCESS", msg, args...)
}

func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
	l.logMessage(WARN, "WARN", msg, args...)
}

func (l *DefaultLogger) Error(msg string, args ...interface{}) {
	l.logMessage(ERROR, "ERROR", msg, args...)
}

func (l *DefaultLogger) logMessage(level LogLevel, levelStr, msg string, args ...interface{}) {
	if l.level > level {
		return
	}

	var formattedMsg string
	if len(args) > 0 {
		formattedMsg = fmt.Sprintf(msg, args...)
	} else {
		formattedMsg = msg
	}

	coloredLevel := l.colorFuncs[level](levelStr)
	l.logger.Printf("[%s] %s", coloredLevel, formattedMsg)
}
