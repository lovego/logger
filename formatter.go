package logger

import (
	"encoding/json"
)

var jsonFormatter jsonFmt
var jsonIndentFormatter jsonIndentFmt

type Formatter interface {
	Format(map[string]interface{}) ([]byte, error)
}

type jsonFmt struct {
}

func (jf jsonFmt) Format(fields map[string]interface{}) ([]byte, error) {
	return json.Marshal(fields)
}

type jsonIndentFmt struct {
}

func (jif jsonIndentFmt) Format(fields map[string]interface{}) ([]byte, error) {
	return json.Marshal(fields)
}
