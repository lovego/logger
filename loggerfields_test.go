package logger

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestSet(t *testing.T) {
	logger := New(nil)
	if logger.Set("key", "value"); !reflect.DeepEqual(logger.fields,
		map[string]interface{}{"key": "value"}) {
		t.Errorf("unexpected logger fields %v", logger.fields)
	}
}

func TestSetMachineName(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.SetMachineName()
	if _, ok := logger.fields["machineName"]; !ok {
		t.Errorf("unexpected logger %v", logger)
	}
	logger.Info("the message")

	hostname, _ := os.Hostname()
	expect := fmt.Sprintf(`"level":"info","machineName":"%s","msg":"the message"}
`, hostname)

	if !strings.HasSuffix(writer.String(), expect) {
		t.Errorf("unexpected output: %s", writer.String())
	} else {
		t.Log(writer.String())
	}
}

func TestSetMachineIP(t *testing.T) {
	logger := New(nil)
	logger.SetMachineIP()
	if _, ok := logger.fields["machineIP"]; !ok {
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
