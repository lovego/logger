package logger

import (
	"context"
	"time"

	"github.com/lovego/tracer"
)

func (l *Logger) Record(
	workFunc func(context.Context) error, recoverFunc func(), fieldsFunc func(*Fields),
) {
	l.RecordWithContext(context.Background(), workFunc, recoverFunc, fieldsFunc)
}

func (l *Logger) RecordWithContext(ctx context.Context,
	workFunc func(context.Context) error, recoverFunc func(), fieldsFunc func(*Fields),
) {
	span := &tracer.Span{At: time.Now()}
	ctx = tracer.Context(ctx, span)
	var err error
	defer func() {
		panicErr := recover()
		if panicErr != nil && recoverFunc != nil {
			recoverFunc()
		}

		f := l.WithSpan(span)
		if fieldsFunc != nil {
			fieldsFunc(f)
		}

		if panicErr != nil {
			f.output(Recover, panicErr, f.data)
		} else if err != nil {
			f.output(Error, err, f.data)
		} else {
			f.output(Info, nil, f.data)
		}
	}()
	err = workFunc(ctx)
}

func (l *Logger) WithSpan(span *tracer.Span) *Fields {
	span.Finish()
	f := l.With("at", span.At).With("duration", span.Duration)
	if len(span.Children) > 0 {
		f = f.With("children", span.Children)
	}
	if len(span.Tags) > 0 {
		f = f.With("tags", span.Tags)
	}
	if len(span.Logs) > 0 {
		f = f.With("logs", span.Logs)
	}
	return f
}
