package logger

import (
	"fmt"
	"io"
	"os"
)

type Logger struct {
	level          Level
	fields         map[string]interface{}
	formatter      Formatter
	writer         io.Writer
	alarmFormatter Formatter
	alarm          Alarm
}

func New(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stderr
	}
	return &Logger{
		level: Info, writer: writer,
		formatter: jsonFormatter, alarmFormatter: jsonIndentFormatter,
	}
}

func (l *Logger) With(key string, value interface{}) *Fields {
	return &Fields{logger: l, data: map[string]interface{}{key: value}}
}

func (l *Logger) Debug(args ...interface{}) bool {
	if len(args) > 0 && l.level >= Debug {
		l.output(Debug, fmt.Sprint(args...), nil)
	}
	return l.level >= Debug
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.level >= Debug {
		l.output(Debug, fmt.Sprintf(format, args...), nil)
	}
}

func (l *Logger) Info(args ...interface{}) bool {
	if len(args) > 0 && l.level >= Info {
		l.output(Info, fmt.Sprint(args...), nil)
	}
	return l.level >= Info
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.level >= Info {
		l.output(Info, fmt.Sprintf(format, args...), nil)
	}
}

func (l *Logger) Error(args ...interface{}) {
	l.output(Error, fmt.Sprint(args...), nil)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.output(Error, fmt.Sprintf(format, args...), nil)
}

func (l *Logger) Recover() {
	if err := recover(); err != nil {
		l.output(Panic, fmt.Sprintf("PANIC: %v", err), nil)
	}
}

func (l *Logger) Panic(args ...interface{}) {
	msg := fmt.Sprint(args...)
	l.output(Panic, fmt.Sprint(args...), nil)
	panic(msg)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.output(Panic, msg, nil)
	panic(msg)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.output(Fatal, fmt.Sprint(args...), nil)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.output(Fatal, fmt.Sprintf(format, args...), nil)
	os.Exit(1)
}
