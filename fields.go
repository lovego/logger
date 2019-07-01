package logger

import (
	"fmt"
	"os"
)

type Fields struct {
	*Logger
	data map[string]interface{}
}

// don't use (level, at, msg, stack, duration) as key, they will be overwritten.
func (f *Fields) With(key string, value interface{}) *Fields {
	f.data[key] = value
	return f
}

func (f *Fields) Debug(args ...interface{}) bool {
	if len(args) > 0 && f.level >= Debug {
		setStackField(f.data, 0, args...)
		f.output(Debug, fmt.Sprint(args...), f.data)
	}
	return f.level >= Debug
}

func (f *Fields) Debugf(format string, args ...interface{}) {
	if f.level >= Debug {
		setStackField(f.data, 0, args...)
		f.output(Debug, fmt.Sprintf(format, args...), f.data)
	}
}

func (f *Fields) Info(args ...interface{}) bool {
	if len(args) > 0 && f.level >= Info {
		setStackField(f.data, 0, args...)
		f.output(Info, fmt.Sprint(args...), f.data)
	}
	return f.level >= Info
}

func (f *Fields) Infof(format string, args ...interface{}) {
	if f.level >= Info {
		setStackField(f.data, 0, args...)
		f.output(Info, fmt.Sprintf(format, args...), f.data)
	}
}

func (f *Fields) Error(args ...interface{}) {
	setStackField(f.data, 4, args...)
	f.output(Error, fmt.Sprint(args...), f.data)
}

func (f *Fields) Errorf(format string, args ...interface{}) {
	setStackField(f.data, 4, args...)
	f.output(Error, fmt.Sprintf(format, args...), f.data)
}

func (f *Fields) Recover() {
	if err := recover(); err != nil {
		setStackField(f.data, 5, err)
		f.output(Recover, fmt.Sprint(err), f.data)
	}
}

func (f *Fields) Panic(args ...interface{}) {
	setStackField(f.data, 4, args...)
	msg := fmt.Sprint(args...)
	f.output(Panic, msg, f.data)
	panic(msg)
}

func (f *Fields) Panicf(format string, args ...interface{}) {
	setStackField(f.data, 4, args...)
	msg := fmt.Sprintf(format, args...)
	f.output(Panic, msg, f.data)
	panic(msg)
}

func (f *Fields) Fatal(args ...interface{}) {
	setStackField(f.data, 4, args...)
	f.output(Fatal, fmt.Sprint(args...), f.data)
	os.Exit(1)
}

func (f *Fields) Fatalf(format string, args ...interface{}) {
	setStackField(f.data, 4, args...)
	f.output(Fatal, fmt.Sprintf(format, args...), f.data)
	os.Exit(1)
}
