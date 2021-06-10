# logger
a logger that integrate with alarm.

[![Build Status](https://github.com/lovego/logger/actions/workflows/go.yml/badge.svg)](https://github.com/lovego/logger/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/lovego/logger/badge.svg?branch=master)](https://coveralls.io/github/lovego/logger)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/logger)](https://goreportcard.com/report/github.com/lovego/logger)
[![Documentation](https://pkg.go.dev/badge/github.com/lovego/logger)](https://pkg.go.dev/github.com/lovego/logger@v0.0.1)

## Install
`$ go get github.com/lovego/logger`

## Usage
```go
logger := New(os.Stdout)

logger.SetLevel(Debug)
logger.Debug("the ", "message")
logger.Debugf("this is %s", "test")

logger.Info("the ", "message")
logger.Infof("this is a %s", "test")

logger.Error("err")
logger.Errorf("test %s", "errorf")

logger.Panic("panic !!")
logger.Panicf("test %s", "panicf")

logger.Fatal("fatal !!")
logger.Fatalf("test %s", "fatalf")

defer logger.Recover()

logger.Record(func(ctx context.Context) error {
  // work to do goes here
  return nil
}, nil, func(f *Fields) {
  f.With("key1", "value1").With("key2", "value2")
})
```

