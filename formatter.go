package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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

	var msg, stack = fields["msg"], fields["stack"]
	if msg != nil && msg != "" {
		buf.WriteString(fmt.Sprintf("%v\n", msg))
		delete(fields, "msg")
		defer func() { fields["msg"] = msg }()
	}
	if stack != nil && stack != "" {
		buf.WriteString(fmt.Sprintf("%v\n", stack))
		delete(fields, "stack")
		defer func() { fields["stack"] = stack }()
	}

	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func PrintJson(v interface{}) {
	data, err := json.MarshalIndent(v, ``, `  `)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(data))
}
