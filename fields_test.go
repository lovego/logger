package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/bouk/monkey"
)

func TestFieldsWith(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Info(`the `, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsDebug(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).SetLevel(Debug).With("key", "value").With("key2", "value2").
		Debug(`the `, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"debug","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsDebugf(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).SetLevel(Debug).With("key", "value").With("key2", "value2").
		Debugf("%s %s", `the`, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"debug","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsInfo(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Info(`the `, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsInfof(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Infof("%s %s", `the`, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsError(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Error(`the `, `message`)
	if !strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"error","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestFieldsError\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsErrorf(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Errorf("%s %s", `the`, `message`)
	if !strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"error","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestFieldsErrorf\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsRecover(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	func() {
		defer New(writer).With("key", "value").With("key2", "value2").Recover()
		panic("the message")
	}()
	if !strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"recover","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestFieldsRecover.func1\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsPanic(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	defer func() {
		err := recover()
		if err != "the message" {
			t.Errorf("unexpected err: %v", err)
		}
		if !strings.Contains(writer.String(),
			`,"key":"value","key2":"value2","level":"panic","msg":"the message",`+
				`"stack":"github.com/lovego/logger.TestFieldsPanic\n\t`) {
			t.Errorf("unexpected output: %s", writer.String())
		}
	}()
	New(writer).With("key", "value").With("key2", "value2").Panic("the message")
}

func TestFieldsPanicf(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	defer func() {
		err := recover()
		if err != "the message" {
			t.Errorf("unexpected err: %v", err)
		}
		if !strings.Contains(writer.String(),
			`,"key":"value","key2":"value2","level":"panic","msg":"the message",`+
				`"stack":"github.com/lovego/logger.TestFieldsPanicf\n\t`) {
			t.Errorf("unexpected output: %s", writer.String())
		}
	}()
	New(writer).With("key", "value").With("key2", "value2").Panicf("%s %s", "the", "message")
}

func TestFieldsFatal(t *testing.T) {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Fatal("the message")
	if exitStatus != 1 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}
	if !strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"fatal","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestFieldsFatal\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestFieldsFatalf(t *testing.T) {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Fatalf("%s %s", "the", "message")
	if exitStatus != 1 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}
	if !strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"fatal","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestFieldsFatalf\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}
