package logger

type Fields struct {
	logger *Logger
	data   map[string]interface{}
}

func (f *Fields) With(key string, value interface{}) *Fields {
	f.data[key] = value
	return f
}
