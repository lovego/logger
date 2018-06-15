package logger

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/lovego/errs"
	"github.com/lovego/tracer"
)

func TestWithSpan(t *testing.T) {
	logger := New(nil)
	var cSpan = []*tracer.Span{{At: time.Now()}}
	var span = &tracer.Span{
		At:       time.Now(),
		Children: cSpan,
		Tags:     map[string]interface{}{"key": "value"},
	}
	if got := logger.WithSpan(span); got.Logger == nil {
		t.Errorf("unexpect got Fields %v", got)
	} else {
		if _, ok := got.data[`at`]; !ok {
			t.Errorf("field not be seted %v", got.data)
		}

		if _, ok := got.data[`duration`]; !ok {
			t.Errorf("field not be seted %v", got.data)
		}

		if _, ok := got.data[`children`]; !ok {
			t.Errorf("field not be seted %v", got.data)
		}

		if _, ok := got.data[`tags`]; !ok {
			t.Errorf("field not be seted %v", got.data)
		}
	}
}

func TestRecord1(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(true, nil, func() {}, func(*Fields) {})
	if writer.String() == `` {
		t.Errorf("test %s", writer.String())
	}
}

func TestRecord2(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(true, func(ctx context.Context) error {
		return errs.New("code", "message")
	}, nil, nil)
	if writer.String() == `` {
		t.Errorf("test %s", writer.String())
	}
}

func TestRecord3(t *testing.T) {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.Record(false, func(ctx context.Context) error {
		return nil
	}, nil, nil)
	if writer.String() == `` {
		t.Errorf("test %s", writer.String())
	}
}
