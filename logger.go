package logger

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"

	"github.com/lovego/errs"
)

type logger struct {
	level     uint8
	fields    map[string]interface{}
	formatter Formatter
	writer    io.Writer
	alarm     Alarm
}

func New(writer io.Writer) *Logger {
	if writer == nil {
		writer = os.Stderr
	}
	return &Logger{writer: writer}
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
	l.doAlarm(fmt.Sprint(args...))
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.doAlarm(fmt.Sprintf(format, args...))
}

func (l *Logger) doAlarm(title string) {
	stack := errs.Stack(4)
	content := l.output(title) + stack
	l.writer.Write([]byte(content))
	title = l.prefix + ` ` + title
	mergeKey := title + "\n" + stack // 根据title和调用栈对报警消息进行合并
	if l.alarm != nil {
		l.alarm.Alarm(title, content, mergeKey)
	}
}

func (l *Logger) Alarm(titleValue, contentValue interface{}) {
	title := fmt.Sprint(titleValue)
	content := fmt.Sprint(contentValue)
	stack := errs.Stack(3)

	content = l.output(title) + content + "\n" + stack
	l.writer.Write([]byte(content))
	title = l.prefix + ` ` + title
	mergeKey := title + "\n" + stack // 根据title和调用栈对报警消息进行合并
	if l.alarm != nil {
		l.alarm.Alarm(title, content, mergeKey)
	}
}

func (l *Logger) Recover() {
	if err := recover(); err != nil {
		l.doAlarm(fmt.Sprintf("PANIC: %v\n", err))
	}
}

func (l *Logger) Fatal(args ...interface{}) {
	l.doExit(fmt.Sprint(args...))
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.doExit(fmt.Sprintf(format, args...))
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.doExit(fmt.Sprintln(args...))
}

func (l *Logger) doExit(title string) {
	stack := errs.Stack(4)
	content := l.output(title) + stack
	l.writer.Write([]byte(content))
	title = l.prefix + ` ` + title
	if l.alarm != nil {
		l.alarm.Send(title, content)
	}
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.doPanic(fmt.Sprint(args...))
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.doPanic(fmt.Sprintf(format, args...))
}

func (l *Logger) Panicln(args ...interface{}) {
	l.doPanic(fmt.Sprintln(args...))
}

func (l *Logger) doPanic(title string) {
	stack := errs.Stack(4)
	titleLine := l.output(title)

	content := titleLine + stack
	if l.writer != os.Stderr {
		l.writer.Write([]byte(content))
	}
	title = l.prefix + ` ` + title
	if l.alarm != nil {
		l.alarm.Send(title, content)
	}

	os.Stderr.Write([]byte(titleLine))
	panic(title)
}
