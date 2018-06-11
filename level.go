package logger

type Level int8

const (
	Fatal Level = iota
	Panic
	Error
	Info
	Debug
)

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
