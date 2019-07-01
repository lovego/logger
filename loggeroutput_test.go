package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func ExampleLogger_output_1() {
	writer := bytes.NewBuffer(nil)
	New(writer).SetPid().output(Info, "message", map[string]interface{}{"key": "value"})
	expect := fmt.Sprintf(`,"key":"value","level":"info","msg":"message","pid":%d}`, os.Getpid())
	fmt.Println(strings.Contains(writer.String(), expect))
	// Output: true
}

func ExampleLogger_output_2() {
	writer := bytes.NewBuffer(nil)
	New(writer).output(Error, "message", map[string]interface{}{"key": "value"})
	fmt.Println(strings.Contains(writer.String(), `,"key":"value","level":"error","msg":"message"}`))
	// Output: true
}

func ExampleLogger_output_3() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.writer = os.Stderr
	logger.output(Panic, "message", map[string]interface{}{"key": "value"})
	fmt.Println(writer.String())
	// Output:
}

func ExampleLogger_getFields_1() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.fields = map[string]interface{}{"key": true, "key1": "value1"}
	got := logger.getFields(Recover, "message", nil)
	fmt.Println(got[`level`], got[`msg`], got[`key`], got[`key1`])
	// Output:
	// recover message true value1
}

func ExampleLogger_doAlarm_1() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	logger := New(writer)
	logger.SetAlarm(alarm)
	logger.doAlarm(Panic, nil)
	fmt.Println(alarm.title, alarm.content)
	// Output: null
}

func ExampleLogger_doAlarm_2() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	var mapIn = make(map[interface{}]interface{})
	logger.doAlarm(Panic, map[string]interface{}{"test": mapIn})
	fmt.Println(strings.Contains(writer.String(),
		`"level":"error","msg":"logger format: json: unsupported type: map[interface {}]interface {}`,
	))
	// Output: true
}

func ExampleLogger_format_1() {
	writer, alarm := bytes.NewBuffer(nil), &testAlarm{}
	logger := New(writer)
	logger.SetAlarm(alarm)
	fmt.Println(string(logger.format(map[string]interface{}{"key": true}, true)))
	// Output:
	// {
	//   "key": true
	// }
}

func ExampleLogger_format_2() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	var testIn = make(map[interface{}]interface{})
	var fields = make(map[string]interface{})
	fields[`key`] = testIn
	fmt.Println(logger.format(fields, true) == nil)
	// Output: true
}

func ExampleLogger_SetLevel() {
	logger := New(nil)
	logger.SetLevel(Panic)
	fmt.Println(logger.level)

	var level Level = 10
	logger.SetLevel(level)
	fmt.Println(logger.level)

	logger.SetLevel(Info)
	fmt.Println(logger.level)

	// Ouput:
	// panic
	// debug
	// info
}

func ExampleLevel_String() {
	var l Level
	l = 10
	fmt.Println(l.String())
	// Ouput: invalid
}
