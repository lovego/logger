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

func ExampleLogger_WithTracer() {
	logger := New(nil)
	var t = &tracer.Tracer{
		At:       time.Now(),
		Children: []*tracer.Tracer{{At: time.Now()}},
		Tags:     map[string]interface{}{"key": "value"},
		Logs:     []string{"the message"},
	}
	got := logger.WithTracer(t)

	expect := &Fields{
		Logger: logger,
		data: map[string]interface{}{
			"at": t.At, "duration": t.Duration,
			"children": t.Children, "tags": t.Tags,
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
	fmt.Println(strings.Contains(writer.String(),
		`"key":"value","level":"recover",`+machineName+`"msg":"the message",`+
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
	fmt.Println(strings.Contains(writer.String(), `,"level":"error",`+machineName+`"msg":"code: message",`+
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
		`,"level":"error",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleLogger_Record_3.func1\n\t`,
	))
	// Output: true
}

func ExampleLogger_Record_4() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(func(ctx context.Context) error {
		tracer.Tag(ctx, "tagKey", "tagValue")
		t := tracer.Start(ctx, "test")
		tracer.Tag(t, "tagK", "tagV")
		defer tracer.Finish(t)
		return nil
	}, nil, nil)
	s := writer.String()
	fmt.Println(strings.Contains(s, `,"level":"info",`+machineName+`"tags":{"tagKey":"tagValue"}}`))
	fmt.Println(strings.Contains(s, `,"children":[{"name":"test","at":`))
	fmt.Println(strings.Contains(s, `,"tags":{"tagK":"tagV"}}],"duration":`))
	// Output:
	// true
	// true
	// true
}
