package logger

import (
	"testing"
	"time"

	"github.com/lovego/tracer"
)

func TestSpanFields(t *testing.T) {
	logger := New(nil)
	var cSpan = []*tracer.Span{{At: time.Now()}}
	var span = &tracer.Span{
		At:       time.Now(),
		Children: cSpan,
		Tags:     map[string]interface{}{"key": "value"},
	}
	if got := logger.spanFields(span); got.Logger == nil {
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

func TestRecord(t *testing.T) {
	logger := New(nil)

}
