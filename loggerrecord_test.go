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
	span.Log("the ", "message")
	got := logger.WithSpan(span)
	expect := &Fields{
		Logger: logger,
		data: map[string]interface{}{
			"at": span.At, "duration": span.Duration,
			"children": span.Children, "tags": span.Tags,
			"logs": []string{"the message"},
		},
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("unexpected got: %+v", got)
	}
}

func TestRecord1(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
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
	logger.Record(func(ctx context.Context) error {
		return errs.New("code", "message")
	}, nil, nil)
	if !strings.Contains(writer.String(), `,"level":"error","msg":"code: message",`+
		`"stack":"github.com/lovego/logger.(*Logger).RecordWithContext\n\t`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestRecord3(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		return errs.Tracef("the message")
	}, nil, nil)
	if !strings.Contains(writer.String(),
		`,"level":"error","msg":"the message",`+
			`"stack":"github.com/lovego/logger.TestRecord3.func1\n\t`,
	) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}

func TestRecord4(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		tracer.Tag(ctx, "tagKey", "tagValue")
		span := tracer.StartSpan(ctx, "test")
		span.Tag("tagK", "tagV")
		defer span.Finish()
		return nil
	}, nil, nil)
	s := writer.String()
	if !strings.Contains(s, `,"children":[{"name":"test","at":`) ||
		!strings.Contains(s, `,"tags":{"tagK":"tagV"}}],"duration":`) ||
		!strings.Contains(s, `,"level":"info","tags":{"tagKey":"tagValue"}}`) {
		t.Errorf("unexpected output: %s", writer.String())
	}
}
