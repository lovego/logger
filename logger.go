package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/lovego/errs"
)

const (
	Fatal Level = iota
	Panic
	Recover
	Error
	Info
	Debug
)

type Level int8
type Logger struct {
	level          Level
	fields         map[string]interface{}
	formatter      Formatter
	alarmFormatter Formatter
	writer         io.Writer
	alarm          Alarm
}

type Alarm interface {
	Send(title, content string)
	Alarm(title, content, mergeKey string)
}

func New(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stderr
	}
	var formatter Formatter
	if writer == os.Stdout || writer == os.Stderr {
		formatter = readableFormatter
	} else {
		formatter = jsonFormatter
	}
	hostname, _ := os.Hostname()

	return &Logger{
		level: Info, writer: writer,
		formatter: formatter, alarmFormatter: readableFormatter,
		fields: map[string]interface{}{"machineName": hostname},
	}
}

// don't use (level, at, msg, stack, duration) as key, they will be overwritten.
func (l *Logger) With(key string, value interface{}) *Fields {
	return &Fields{Logger: l, data: map[string]interface{}{key: value}}
}

func (l *Logger) Debug(args ...interface{}) bool {
	if len(args) > 0 && l.level >= Debug {
		l.output(Debug, fmt.Sprint(args...), getStackField(0, args...))
	}
	return l.level >= Debug
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= Debug {
		l.output(Debug, fmt.Sprintf(format, args...), getStackField(0, args...))
	}
}

func (l *Logger) Info(args ...interface{}) bool {
	if len(args) > 0 && l.level >= Info {
		l.output(Info, fmt.Sprint(args...), getStackField(0, args...))
	}
	return l.level >= Info
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= Info {
		l.output(Info, fmt.Sprintf(format, args...), getStackField(0, args...))
	}
}

func (l *Logger) Error(args ...interface{}) {
	l.output(Error, fmt.Sprint(args...), getStackField(4, args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.output(Error, fmt.Sprintf(format, args...), getStackField(4, args...))
}

func (l *Logger) Recover() {
	if err := recover(); err != nil {
		l.output(Recover, fmt.Sprint(err), getStackField(4+errs.PanicStackDepth(), err))
	}
}

func (l *Logger) Panic(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.output(Panic, fmt.Sprint(args...), getStackField(4, args...))
	panic(msg)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.output(Panic, msg, getStackField(4, args...))
	panic(msg)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.output(Fatal, fmt.Sprint(args...), getStackField(4, args...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.output(Fatal, fmt.Sprintf(format, args...), getStackField(4, args...))
	os.Exit(1)
}
