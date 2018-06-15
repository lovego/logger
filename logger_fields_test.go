package logger

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	logger := New(nil)
	if logger.Set("key", "value"); !reflect.DeepEqual(logger.fields,
		map[string]interface{}{"key": "value"}) {
		t.Errorf("unexpected logger fields %v", logger.fields)
	}
}

func TestSethostname(t *testing.T) {
	logger := New(nil)
	logger.SetHostname()
	if _, ok := logger.fields["hostname"]; !ok {
		t.Errorf("unexpected logger %v", logger)
	}
}
func TestSetIP(t *testing.T) {
	logger := New(nil)
	logger.SetIP()
	if _, ok := logger.fields["ip"]; !ok {
		t.Errorf("unexpected logger %v", logger)
	}
}

func TestSetPid(t *testing.T) {
	logger := New(nil)
	logger.SetPid()
	if _, ok := logger.fields["pid"]; !ok {
		t.Errorf("unexpected logger %v", logger)
	}
}
