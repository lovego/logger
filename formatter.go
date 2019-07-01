package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var jsonFormatter jsonFmt
var readableFormatter readableFmt

type Formatter interface {
	Format(map[string]interface{}) ([]byte, error)
}

type jsonFmt struct {
}

func (jf jsonFmt) Format(fields map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(fields); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type readableFmt struct {
}

func (rf readableFmt) Format(fields map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	msg, stack := fields["msg"], fields["stack"]
	if msg != nil || stack != nil {
		delete(fields, "msg")
		delete(fields, "stack")
		err := encoder.Encode(fields)
		if err != nil {
			return nil, err
		}
		fields["msg"], fields["stack"] = msg, stack
		return append([]byte(fmt.Sprintf("%v\n%v\n", msg, stack)), buf.Bytes()...), nil
	}
	return json.MarshalIndent(fields, "", "  ")
}
