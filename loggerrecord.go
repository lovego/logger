package logger

import (
	"context"

	"github.com/lovego/errs"
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
	ctx = tracer.Start(ctx, ``)
	var err error
	defer func() {
		panicErr := recover()
		if panicErr != nil && recoverFunc != nil {
			recoverFunc()
		}

		tracer.Finish(ctx)
		f := l.WithTracer(tracer.Get(ctx))
		if fieldsFunc != nil {
			fieldsFunc(f)
		}

		if panicErr != nil {
			setStackField(f.data, 4+errs.PanicStackDepth(), panicErr)
			f.output(Recover, panicErr, f.data)
		} else if err != nil {
			setStackField(f.data, 4, err)
			f.output(Error, err, f.data)
		} else {
			f.output(Info, nil, f.data)
		}
	}()
	err = workFunc(ctx)
}

func (l *Logger) WithTracer(t *tracer.Tracer) *Fields {
	if t == nil {
		return &Fields{Logger: l, data: map[string]interface{}{}}
	}

	f := l.With("at", t.At).With("duration", t.Duration)
	if len(t.Children) > 0 {
		f = f.With("children", t.Children)
	}
	if len(t.Tags) > 0 {
		f = f.With("tags", t.Tags)
	}
	if len(t.Logs) > 0 {
		f = f.With("logs", t.Logs)
	}
	return f
}
