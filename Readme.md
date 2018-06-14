# logger
a logger that integrate with alarm.

[![Build Status](https://travis-ci.org/lovego/logger.svg?branch=new_json)](https://travis-ci.org/lovego/logger)
[![Coverage Status](https://coveralls.io/repos/github/lovego/logger/badge.svg?branch=new_json)](https://coveralls.io/github/lovego/logger?branch=new_json)
[![Go Report Card](https://goreportcard.com/badge/github.com/lovego/logger)](https://goreportcard.com/report/github.com/lovego/logger)
[![GoDoc](https://godoc.org/github.com/lovego/logger?status.svg)](https://godoc.org/github.com/lovego/logger)

## Install
`$ go get github.com/lovego/logger`

## Usage
```go
writer := bytes.NewBuffer(nil)
  logger := New(writer)

  logger.SetLevel(Debug)
  logger.Debug("the", "message")
  fmt.Println(writer.String()) // {"at":"2018-06-14T18:03:18.667846674+08:00","level":"debug","msg":"themessage"}
  logger.Debugf("this is %s", "test")
  fmt.Println(writer.String()) // {"at":"2018-06-14T18:08:07.484039501+08:00","level":"debug","msg":"this is test"}

  logger.Info("the ", "message")
  fmt.Println(writer) // {"at":"2018-06-14T18:12:41.734621203+08:00","level":"info","msg":"the message"}
  logger.Infof("this is a %s", "test")
  fmt.Println(writer) // {"at":"2018-06-14T18:14:00.476463434+08:00","level":"info","msg":"this is a test"}

  logger.Error("err")
  fmt.Println(writer) // {"at":"2018-06-14T18:14:52.04227535+08:00","level":"error","msg":"err","stack":"github.com/lovego/logger.TestReadme\n\t/Users/chenyun/go/src/github.com/lovego/logger/logger_output_test.go:132 (0x112040b)\ntesting.tRunner\n\t/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746 (0x10bf75f)\n"}
  logger.Errorf("test %s", "errorf")
  fmt.Println(writer) // {"at":"2018-06-14T18:16:01.930189701+08:00","level":"error","msg":"test errorf","stack":"github.com/lovego/logger.TestReadme\n\t/Users/chenyun/go/src/github.com/lovego/logger/logger_output_test.go:134 (0x1120420)\ntesting.tRunner\n\t/usr/local/Cellar/go/1.9.2/libexec/src/testing/testing.go:746 (0x10bf75f)\n"}

  logger.Panic("panic !!")
  fmt.Println(writer) // panic: panic !! [recovered]
  logger.Panicf("test %s", "panicf")
  fmt.Println(writer) // panic: test panicf [recovered]

  logger.Fatal("fatal !!")
  fmt.Println(writer) // exit status 1
  logger.Fatalf("test %s", "fatalf")
  fmt.Println(writer) // exit status 1

  logger.Recover()
  fmt.Println(writer) // here is no message back

```

## Documentation
[https://godoc.org/github.com/lovego/logger](https://godoc.org/github.com/lovego/logger)
