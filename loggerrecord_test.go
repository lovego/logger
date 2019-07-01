package logger

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/lovego/errs"
	"github.com/lovego/tracer"
)

func ExampleLogger_WithSpan() {
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
	fmt.Println(reflect.DeepEqual(got, expect))
	// Output: true
}

func ExampleLogger_Record_1() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		panic("the message")
	}, func() {}, func(f *Fields) {
		f.With("key", "value")
	})
	fmt.println(writer.String())
	fmt.Println(strings.Contains(writer.String(),
		`"key":"value","level":"recover","msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleLogger_Record_1.func1\n\t`,
	))
	// Output: true
}

func ExampleLogger_Record_2() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		return errs.New("code", "message")
	}, nil, nil)
	fmt.Println(strings.Contains(writer.String(), `,"level":"error","msg":"code: message",`+
		`"stack":"github.com/lovego/logger.(*Logger).RecordWithContext\n\t`,
	))
	// Output: true
}

func ExampleLogger_Record_3() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		return errs.Tracef("the message")
	}, nil, nil)
	fmt.Println(strings.Contains(writer.String(),
		`,"level":"error","msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleLogger_Record_3.func1\n\t`,
	))
	// Output: true
}

func ExampleLogger_Record_4() {
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
	fmt.Println(strings.Contains(s, `,"children":[{"name":"test","at":`))
	fmt.Println(strings.Contains(s, `,"tags":{"tagK":"tagV"}}],"duration":`))
	fmt.Println(strings.Contains(s, `,"level":"info","tags":{"tagKey":"tagValue"}}`))
	// Output:
	// true
	// true
	// true
}
