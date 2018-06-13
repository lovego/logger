package logger

type Interface interface {
	With(key string, value interface{}) Logger

	Fatal(args ...interface{}) bool
	Fatalf(format string, args ...interface{})
	Panic(args ...interface{}) bool
	Panicf(format string, args ...interface{})
	Error(args ...interface{}) bool
	Errorf(format string, args ...interface{})
	Info(args ...interface{}) bool
	Infof(format string, args ...interface{})
	Debug(args ...interface{}) bool
	Debugf(format string, args ...interface{})
}
