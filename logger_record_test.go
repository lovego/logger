package logger

import (
	"bytes"
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lovego/errs"
	"github.com/lovego/tracer"
)

func TestWithSpan(t *testing.T) {
	logger := New(nil)
	var span = &tracer.Span{
		At:       time.Now(),
		Children: []*tracer.Span{{At: time.Now()}},
		Tags:     map[string]interface{}{"key": "value"},
	}
	got := logger.WithSpan(span)
	expect := &Fields{
		Logger: logger,
		data: map[string]interface{}{
			"at": span.At, "duration": span.Duration,
			"children": span.Children, "tags": span.Tags,
		},
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("unexpected got: %+v", got)
	}
}

func TestRecord1(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(true, func(ctx context.Context) error {
		panic("the message")
	}, func() {}, func(f *Fields) {
		f.With("key", "value")
	})
	if !strings.Contains(writer.String(),
		`"key":"value","level":"recover","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestRecord1.func1\n\t`,
	) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestRecord2(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(true, func(ctx context.Context) error {
		return errs.New("code", "message")
	}, nil, nil)
	if !strings.Contains(writer.String(),
		`,"level":"error","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestRecord2.func1\n\t`,
	) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestRecord3(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(false, func(ctx context.Context) error {
		return nil
	}, nil, nil)
	if !strings.Contains(writer.String(), `,"level":"info"}`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}
