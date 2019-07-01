package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"bou.ke/monkey"
)

type testAlarm struct {
	title, content, mergeKey string
}

func (a *testAlarm) Send(title, content string) {
	a.title, a.content = title, content
}

func (a *testAlarm) Alarm(title, content, mergeKey string) {
	a.title, a.content, a.mergeKey = title, content, mergeKey
}

func ExampleNew() {
	log := New(nil)
	fmt.Println(log.writer == os.Stderr)
	// Output: true
}

func ExampleWith() {
	writer := bytes.NewBuffer(nil)
	log := New(writer)
	log.With("key", "value").Info(`the `, `message`)
	fmt.Println(strings.HasSuffix(
		writer.String(), `,"key":"value","level":"info","msg":"the message"}
`))
	// Output: true
}

func ExampleDebug() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.SetLevel(Debug)
	log.Debug(`the `, `message`)
	fmt.Println(strings.HasSuffix(writer.String(), `"level":"debug","msg":"the message"}
`))
	fmt.Println(*alarm == (testAlarm{}))
	// Output:
	// true
	// true
}

func ExampleDebugf() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.SetLevel(Debug)
	log.Debugf("%s %s", `the`, `message`)
	fmt.Println(strings.HasSuffix(writer.String(), `"level":"debug","msg":"the message"}
`))
	fmt.Println(*alarm == (testAlarm{}))
	// Output:
	// true
	// true
}

func ExampleInfo() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Info(`the `, `message`)
	fmt.Println(strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}
`))
	fmt.Println(*alarm == (testAlarm{}))
	// Output:
	// true
	// true
}

func ExampleInfof() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Infof("%s %s", `the`, `message`)
	fmt.Println(strings.HasSuffix(writer.String(), `"level":"info","msg":"the message"}
`))
	fmt.Println(*alarm == (testAlarm{}))
	// Output:
	// true
	// true
}

func ExampleError() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Error(`the `, `message`)
	fmt.Println(strings.Contains(writer.String(),
		`"level":"error","msg":"the message","stack":"github.com/lovego/logger.ExampleError\n\t`,
	))
	fmt.Println(alarm.title == "the message")
	fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExampleError
`))
	fmt.Println(strings.Contains(alarm.content, `"level": "error"`))
	// Output:
	// true
	// true
	// true
	// true
}

func ExampleErrorf() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Errorf("%s %s", `the`, `message`)
	fmt.Println(strings.Contains(writer.String(),
		`"level":"error","msg":"the message","stack":"github.com/lovego/logger.ExampleErrorf\n\t`,
	))
	fmt.Println(alarm.title == "the message")
	fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExampleErrorf
`))
	fmt.Println(strings.Contains(alarm.content, `"level": "error"`))
	// Output:
	// true
	// true
	// true
	// true
}

func ExampleRecover() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	func() {
		defer log.Recover()
		panic("the message")
	}()
	fmt.Println(strings.Contains(writer.String(), `"level":"recover","msg":"the message",`+
		`"stack":"github.com/lovego/logger.ExampleRecover.func1\n\t`,
	))
	fmt.Println(alarm.title == "the message")
	fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExampleRecover.func1
`))
	fmt.Println(strings.Contains(alarm.content, `"level": "recover"`))
	// Output:
	// true
	// true
	// true
	// true
}

func ExamplePanic() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	defer func() {
		err := recover()
		if err != "the message" {
			fmt.Printf("unexpected err: %v", err)
		}
		fmt.Println(strings.Contains(writer.String(),
			`"level":"panic","msg":"the message","stack":"github.com/lovego/logger.ExamplePanic\n\t`,
		))
		fmt.Println(alarm.title == "the message")
		fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExamplePanic
`))
		fmt.Println(strings.Contains(alarm.content, `"level": "panic"`))
	}()
	log.Panic("the message")
	// Output:
	// true
	// true
	// true
	// true
}

func ExamplePanicf() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	defer func() {
		err := recover()
		if err != "the message" {
			fmt.Printf("unexpected err: %v", err)
		}
		fmt.Println(strings.Contains(writer.String(),
			`"level":"panic","msg":"the message","stack":"github.com/lovego/logger.ExamplePanicf\n\t`,
		))
		fmt.Println(alarm.title == "the message")
		fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExamplePanicf
`))
		fmt.Println(strings.Contains(alarm.content, `"level": "panic"`))
	}()
	log.Panicf("%s %s", "the", "message")
	// Output:
	// true
	// true
	// true
	// true
}

func ExampleFatal() {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Fatal("the message")
	if exitStatus != 1 {
		fmt.Printf("unexpected exit status: %d", exitStatus)
	}
	fmt.Println(strings.Contains(writer.String(),
		`"level":"fatal","msg":"the message","stack":"github.com/lovego/logger.ExampleFatal\n\t`,
	))
	fmt.Println(alarm.title == "the message")
	fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExampleFatal
`))
	fmt.Println(strings.Contains(alarm.content, `"level": "fatal"`))
	// Output:
	// true
	// true
	// true
	// true
}

func ExampleFatalf() {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	log := New(writer)
	log.SetAlarm(alarm)
	log.Fatalf("%s %s", "the", "message")
	if exitStatus != 1 {
		fmt.Printf("unexpected exit status: %d", exitStatus)
	}
	fmt.Println(strings.Contains(writer.String(),
		`"level":"fatal","msg":"the message","stack":"github.com/lovego/logger.ExampleFatalf\n\t`,
	))
	fmt.Println(alarm.title == "the message")
	fmt.Println(strings.Contains(alarm.content, `the message
github.com/lovego/logger.ExampleFatalf
`))
	fmt.Println(strings.Contains(alarm.content, `"level": "fatal"`))
	// Output:
	// true
	// true
	// true
	// true
}
