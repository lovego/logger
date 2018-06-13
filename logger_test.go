package logger

import (
	"bytes"
	"strings"
	"testing"
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

func TestInfo(t *testing.T) {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Info(`the `, `message`)
	if !strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}`) {
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
	if !strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}`) {
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
	if alarm.title != "the message" || !strings.Contains(alarm.content, `,
"level": "error",
"msg": \"the message\",\n\"stack\": \"github.com/lovego/logger.TestError\\n\\t`
  `  {
		t.Errorf("unexpected alarm: %#v", *alarm)
	}
}
