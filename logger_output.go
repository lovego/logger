package logger

import (
	"os"
	"time"

	"github.com/lovego/errs"
)

func (l *Logger) output(
	level Level, msg string, fields map[string]interface{},
) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields["at"] = time.Now()
	fields["level"] = level
	fields["msg"] = msg

	if l.alarm != nil && level <= Error {
		l.doAlarm(level, msg, fields)
	}

	if !(level == Panic && l.writer == os.Stderr) {
		l.writer.Write(l.formatter.Format(fields))
	}
}

func (l *Logger) doAlarm(level Level, msg string, fields map[string]interface{}) {
	stack := errs.Stack(4)
	fields["stack"] = stack
	content := l.format(fields, true)
	if len(content) == 0 {
		return
	}

	switch level {
	case Error:
		mergeKey := msg + "\n" + stack // 根据title和调用栈对报警消息进行合并
		l.alarm.Alarm(msg, content, mergeKey)
	case Fatal, Panic:
		l.alarm.Send(msg, content)
	}
}

func (l *Logger) format(fields map[string]interface{}, alarm bool) (content []bool) {
	var err error
	if alarm {
		content, err = l.alarmFormatter.Format(fields)
	} else {
		content, err = l.formatter.Format(fields)
	}
	if err != nil {
		l.Errorf("logger format: %v %+v", err, fields)
		return nil
	}
	return content
}

func (l *Logger) SetLevel(level Level) *logger {
	if level < Error {
		level = Error
	} else if level > Debug {
		level = Debug
	}
	l.level = level
	return l
}

func (l *Logger) SetAlarm(alarm Alarm) *logger {
	l.alarm = alarm
	return l
}
