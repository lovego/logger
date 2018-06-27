package logger

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/bouk/monkey"
)

type testAlarm struct {
	title, content, mergeKey string
}

func (a *testAlarm) Send(title, content string) {
	a.title, a.content = title, content
}

func (a *testAlarm) Alarm(title, content, mergeKey string) {
	a.title, a.content, a.mergeKey = title, content, mergeKey
}

func TestNew(t *testing.T) {
	log := New(nil)
	if log.writer != os.Stderr {
		t.Errorf("unexpected writer: %s", log.writer)
	}
}

func TestWith(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	log := New(writer)
	log.With("key", "value").Info(`the `, `message`)
	if !strings.HasSuffix(writer.String(),
		`,"key":"value","level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	} else {
		t.Log(writer.String())
	}
}

func TestDebug(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.SetLevel(Debug)
	log.Debug(`the `, `message`)
	if !strings.HasSuffix(writer.String(), `"level":"debug","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if *alarm != (testAlarm{}) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestDebugf(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.SetLevel(Debug)
	log.Debugf("%s %s", `the`, `message`)
	if !strings.HasSuffix(writer.String(), `"level":"debug","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if *alarm != (testAlarm{}) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestInfo(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Info(`the `, `message`)
	if !strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if *alarm != (testAlarm{}) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestInfof(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Infof("%s %s", `the`, `message`)
	if !strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}
`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if *alarm != (testAlarm{}) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestError(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Error(`the `, `message`)
	if !strings.Contains(writer.String(),
		`"level":"error","msg":"the message","stack":"github.com/lovego/logger.TestError\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestError
`) || !strings.Contains(alarm.content, `"level": "error"`) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestErrorf(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Errorf("%s %s", `the`, `message`)
	if !strings.Contains(writer.String(),
		`"level":"error","msg":"the message","stack":"github.com/lovego/logger.TestErrorf\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestErrorf
`) || !strings.Contains(alarm.content, `"level": "error"`) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestRecover(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	func() {
		defer log.Recover()
		panic("the message")
	}()
	if !strings.Contains(writer.String(),
		`"level":"recover","msg":"the message","stack":"github.com/lovego/logger.TestRecover.func1\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestRecover.func1
`) || !strings.Contains(alarm.content, `"level": "recover"`) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestPanic(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	defer func() {
		err := recover()
		if err != "the message" {
			t.Errorf("unexpected err: %v", err)
		}
		if !strings.Contains(writer.String(),
			`"level":"panic","msg":"the message","stack":"github.com/lovego/logger.TestPanic\n\t`) {
			t.Errorf("unexpected output: %s", writer.String())
		}
		if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestPanic
`) || !strings.Contains(alarm.content, `"level": "panic"`) {
			t.Errorf("unexpected alarm: %#v", *alarm)
		}
	}()
	log.Panic("the message")
}

func TestPanicf(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	defer func() {
		err := recover()
		if err != "the message" {
			t.Errorf("unexpected err: %v", err)
		}
		if !strings.Contains(writer.String(),
			`"level":"panic","msg":"the message","stack":"github.com/lovego/logger.TestPanicf\n\t`) {
			t.Errorf("unexpected output: %s", writer.String())
		}
		if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestPanicf
`) || !strings.Contains(alarm.content, `"level": "panic"`) {
			t.Errorf("unexpected alarm: %#v", *alarm)
		}
	}()
	log.Panicf("%s %s", "the", "message")
}

func TestFatal(t *testing.T) {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Fatal("the message")
	if exitStatus != 1 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}
	if !strings.Contains(writer.String(),
		`"level":"fatal","msg":"the message","stack":"github.com/lovego/logger.TestFatal\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestFatal
`) || !strings.Contains(alarm.content, `"level": "fatal"`) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}

func TestFatalf(t *testing.T) {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Fatalf("%s %s", "the", "message")
	if exitStatus != 1 {
		t.Errorf("unexpected exit status: %d", exitStatus)
	}
	if !strings.Contains(writer.String(),
		`"level":"fatal","msg":"the message","stack":"github.com/lovego/logger.TestFatalf\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
	if alarm.title != "the message" || !strings.Contains(alarm.content, `the message
github.com/lovego/logger.TestFatalf
`) || !strings.Contains(alarm.content, `"level": "fatal"`) {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}
