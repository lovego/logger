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
	var buf = &bytes.Buffer{}
	defer formatField(buf, fields, "msg")()
	defer formatField(buf, fields, "data")()
	defer formatField(buf, fields, "stack")()

	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(fields); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func formatField(b *bytes.Buffer, fields map[string]interface{}, fieldName string) func() {
	var fieldValue = fields[fieldName]
	if fieldValue != nil && fieldValue != "" {
		b.WriteString(fmt.Sprintf("%v\n", fieldValue))
		delete(fields, fieldName)
		return func() { fields[fieldName] = fieldValue }
	}
	return func() {}
}

func PrintJson(v interface{}) {
	data, err := json.MarshalIndent(v, ``, `  `)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(data))
}
