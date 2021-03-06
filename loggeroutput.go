package logger

import (
	"fmt"
	"os"
	"time"
)

func (l *Logger) output(
	level Level, msg interface{}, fields map[string]interface{},
) {
	fields = l.getFields(level, msg, fields)

	if level <= Error && l.alarm != nil {
		l.doAlarm(level, fields)
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
	level Level, msg interface{}, fields map[string]interface{},
) map[string]interface{} {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields["level"] = level.String()
	if fields["at"] == nil {
		fields["at"] = time.Now()
	}
	if msg != nil && msg != "" {
		fields["msg"] = fmt.Sprint(msg)
	}

	return fields
}

func (l *Logger) doAlarm(level Level, fields map[string]interface{}) {
	content := l.format(fields, true)
	if len(content) == 0 {
		return
	}
	msg, _ := fields["msg"].(string)
	switch level {
	case Recover, Error:
		stack, _ := fields["stack"].(string)
		l.alarm.Alarm(msg, string(content), msg+"\n"+stack) // 根据msg和stack对报警消息进行合并
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
	case Recover:
		return "recover"
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
