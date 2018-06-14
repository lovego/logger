package logger

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	logger := New(nil)
	if got := logger.Set("key",
		"value"); !reflect.DeepEqual(got.fields, map[string]interface{}{"key": "value"}) {
		t.Errorf("unexpect got fields %v", got.fields)
	}
}

func TestSethostname(t *testing.T) {
	logger := New(nil)
	got := logger.SetHostname()
	if _, ok := got.fields["hostname"]; !ok {
		t.Errorf("test %v", got)
	}
}
func TestSetIP(t *testing.T) {
	logger := New(nil)
	got := logger.SetIP()
	if _, ok := got.fields["ip"]; !ok {
		t.Errorf("test %v", got)
	}
}

func TestSetPid(t *testing.T) {
	logger := New(nil)
	got := logger.SetPid()
	if _, ok := got.fields["pid"]; !ok {
		t.Errorf("test %v", got)
	}
}
