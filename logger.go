package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lovego/tracer"
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
	writer         io.Writer
	alarmFormatter Formatter
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
	return &Logger{
		level: Info, writer: writer,
		formatter: jsonFormatter, alarmFormatter: jsonIndentFormatter,
	}
}

func (l *Logger) With(key string, value interface{}) *Fields {
	return &Fields{Logger: l, data: map[string]interface{}{key: value}}
}

func (l *Logger) Record(debug bool, fun func(ctx context.Context) error) {
	span := &tracer.Span{At: time.Now()}
	if debug {
		span.SetDebug(true)
	}
	var err error
	defer func() {
		f := l.spanFields(span)
		if ret := recover(); ret != nil {
			f.output(Recover, fmt.Sprint(ret), nil)
		} else if err == nil {
			f.output(Info, "", nil)
		} else {
			f.output(Error, err.Error(), nil)
		}
	}()
	err = fun(context.Background())
}

func (l *Logger) spanFields(span *tracer.Span) *Fields {
	span.Finish()
	f := l.With("at", span.At).With("duration", span.Duration)
	if len(span.Children) > 0 {
		f = f.With("children", span.Children)
	}
	if len(span.Tags) > 0 {
		f = f.With("tags", span.Tags)
	}
	return f
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
		l.output(Recover, fmt.Sprint(err), nil)
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
