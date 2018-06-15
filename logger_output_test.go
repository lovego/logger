package logger

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
	// "time"
)

func TestOutput1(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).output(Info, "message", map[string]interface{}{"key": "value"})
	if !strings.Contains(writer.String(),
		`,"key":"value","level":"info","msg":"message"}`) {
		t.Errorf("unexpect output: %s", writer.String())
	}
}

func TestOutput2(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).output(Error, "message", map[string]interface{}{"key": "value"})
	if !strings.Contains(writer.String(),
		`,"key":"value","level":"error","msg":"message","stack":"testing.tRunner\n\t`) {
		t.Errorf("unexpect output: %s", writer.String())
	}
}

func TestOutput3(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.writer = os.Stderr
	logger.output(Panic, "message", map[string]interface{}{"key": "value"})
	if strings.Contains(writer.String(),
		`,"key":"value","level":"panic","msg":"message","stack":"testing.tRunner\n\t`) {
		t.Errorf("unexpect output: %s", writer.String())
	}
}

func TestGetFields(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.fields = map[string]interface{}{"key": true, "key1": "value1"}
	if got := logger.getFields(Recover, "message", nil); got[`level`] != `recover` ||
		got[`msg`] != `message` || got[`key`] != true || got[`key1`] != `value1` {
		t.Errorf("unexpect got %v", got)
	}
}

func TestDoAlarm1(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	logger := New(writer)
	logger.SetAlarm(alarm)
	logger.doAlarm(Panic, nil)
	if alarm.title != `` || alarm.content != `null` {
		t.Errorf("unexpect alarm %v", alarm)
	}
}

func TestDoAlarm2(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	var mapIn = make(map[interface{}]interface{})
	logger.doAlarm(Panic, map[string]interface{}{"test": mapIn})
	if !strings.Contains(writer.String(),
		`"level":"error","msg":"logger format: json: unsupported type: map[interface {}]interface {}`) {
		t.Errorf("unexpect writer %s", writer.String())
	}
}

func TestFormat1(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	logger := New(writer)
	logger.SetAlarm(alarm)
	var expectMap = map[string]interface{}{"key": true}
	expect, err := json.MarshalIndent(expectMap, ``, ``)
	if err != nil {
		t.Errorf("unexpect marshal err %v", err)
	}
	if got := logger.format(map[string]interface{}{"key": true},
		true); string(got) != string(expect) {
		t.Errorf("unexpect got %v,string %s", got, string(got))
	}
}

func TestFormat2(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	var testIn = make(map[interface{}]interface{})
	var fields = make(map[string]interface{})
	fields[`key`] = testIn
	content := logger.format(fields, true)
	if content != nil {
		t.Errorf("unexpect content %s", string(content))
	}
}

func TestSetLevel(t *testing.T) {
	logger := New(nil)
	if got := logger.SetLevel(Panic); got.level != Error {
		t.Errorf("unexpct level %d", logger.level)
	}

	var level Level = 10
	if got := logger.SetLevel(level); got.level != Debug {
		t.Errorf("unexpect level %d", logger.level)
	}

	if got := logger.SetLevel(Info); got.level != Info {
		t.Errorf("unexpect level %d", got.level)
	}
}

func TestString(t *testing.T) {
	var l Level
	l = 10
	if got := l.String(); got != `invalid` {
		t.Errorf("unexpect got %s", got)
	}
}
