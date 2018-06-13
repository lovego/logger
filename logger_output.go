package logger

import (
	"os"
	"time"

	"github.com/lovego/errs"
)

func (l *Logger) output(
	level Level, msg string, fields map[string]interface{},
) {
	fields = l.getFields(level, msg, fields)

	if level <= Error && l.alarm != nil {
		l.doAlarm(level, msg, fields)
	}

	if level == Panic && l.writer == os.Stderr {
		return
	}
	content := l.format(fields, false)
	if len(content) > 0 {
		l.writer.Write(content)
	}
}

func (l *Logger) getFields(
	level Level, msg string, fields map[string]interface{},
) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields["at"] = time.Now()
	fields["level"] = level.String()
	fields["msg"] = msg

	if level <= Error {
		fields["stack"] = errs.Stack(5)
	}
	return fields
}

func (l *Logger) doAlarm(level Level, msg string, fields map[string]interface{}) {
	content := l.format(fields, true)
	if len(content) == 0 {
		return
	}

	switch level {
	case Error:
		mergeKey := msg + "\n" + fields["stack"].(string) // 根据msg和stack对报警消息进行合并
		l.alarm.Alarm(msg, string(content), mergeKey)
	case Fatal, Panic:
		l.alarm.Send(msg, string(content))
	}
}

func (l *Logger) format(fields map[string]interface{}, alarm bool) (content []byte) {
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

func (l *Logger) SetLevel(level Level) *Logger {
	if level < Error {
		level = Error
	} else if level > Debug {
		level = Debug
	}
	l.level = level
	return l
}

func (l *Logger) SetAlarm(alarm Alarm) *Logger {
	l.alarm = alarm
	return l
}

func (l Level) String() string {
	switch l {
	case Fatal:
		return "fatal"
	case Panic:
		return "panic"
	case Error:
		return "error"
	case Info:
		return "info"
	case Debug:
		return "debug"
	default:
		return "invalid"
	}
}
