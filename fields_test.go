package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"bou.ke/monkey"
)

var machineName = getMachineName()

func getMachineName() string {
	name, _ := os.Hostname()
	return fmt.Sprintf(`"machineName":"%s",`, name)
}
func ExampleFields_With() {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Info(`the `, `message`)
	str := writer.String()
	fmt.Println(strings.HasSuffix(
		str, `,"key":"value","key2":"value2","level":"info",`+machineName+`"msg":"the message"}
`))
	// Output: true
}

func ExampleFields_Debug() {
	writer := bytes.NewBuffer(nil)
	New(writer).SetLevel(Debug).With("key", "value").With("key2", "value2").Debug(`the `, `message`)
	fmt.Println(strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"debug",`+machineName+`"msg":"the message"}
`))
	// Output: true
}

func ExampleFields_Debugf() {
	writer := bytes.NewBuffer(nil)
	New(writer).SetLevel(Debug).With("key", "value").With("key2", "value2").
		Debugf("%s %s", `the`, `message`)
	fmt.Println(strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"debug",`+machineName+`"msg":"the message"}
`))
	// Output: true
}

func ExampleFields_Info() {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Info(`the `, `message`)
	fmt.Println(strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"info",`+machineName+`"msg":"the message"}
`))
	// Output: true
}

func ExampleFields_Infof() {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Infof("%s %s", `the`, `message`)
	fmt.Println(strings.HasSuffix(writer.String(),
		`,"key":"value","key2":"value2","level":"info",`+machineName+`"msg":"the message"}
`))
	// Output: true
}

func ExampleFields_Error() {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Error(`the `, `message`)
	fmt.Println(strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"error",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleFields_Error\n\t`,
	))
	// Output: true
}

func ExampleFields_Errorf() {
	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Errorf("%s %s", `the`, `message`)
	fmt.Println(strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"error",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleFields_Errorf\n\t`,
	))
	// Output: true
}

func ExampleFields_Recover() {
	writer := bytes.NewBuffer(nil)
	func() {
		defer New(writer).With("key", "value").With("key2", "value2").Recover()
		panic("the message")
	}()
	fmt.Println(strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"recover",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleFields_Recover.func1\n\t`,
	))
	// Output: true
}

func ExampleFields_Panic() {
	writer := bytes.NewBuffer(nil)
	defer func() {
		err := recover()
		if err != "the message" {
			fmt.Printf("unexpected err: %v", err)
			return
		}
		fmt.Println(strings.Contains(writer.String(),
			`,"key":"value","key2":"value2","level":"panic",`+machineName+`"msg":"the message",`+
				`"stack":"github.com/lovego/logger.ExampleFields_Panic\n\t`,
		))
	}()
	New(writer).With("key", "value").With("key2", "value2").Panic("the message")
	// Output: true
}

func ExampleFields_Panicf() {
	writer := bytes.NewBuffer(nil)
	defer func() {
		err := recover()
		if err != "the message" {
			fmt.Printf("unexpected err: %v", err)
			return
		}
		fmt.Println(strings.Contains(writer.String(),
			`,"key":"value","key2":"value2","level":"panic",`+machineName+`"msg":"the message",`+
				`"stack":"github.com/lovego/logger.ExampleFields_Panicf\n\t`,
		))
	}()
	New(writer).With("key", "value").With("key2", "value2").Panicf("%s %s", "the", "message")
	// Output: true
}

func ExampleFields_Fatal() {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Fatal("the message")
	if exitStatus != 1 {
		fmt.Printf("unexpected exit status: %d", exitStatus)
		return
	}
	fmt.Println(strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"fatal",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleFields_Fatal\n\t`,
	))
	// Output: true
}

func ExampleFields_Fatalf() {
	var exitStatus int
	patch := monkey.Patch(os.Exit, func(status int) {
		exitStatus = status
	})
	defer patch.Unpatch()

	writer := bytes.NewBuffer(nil)
	New(writer).With("key", "value").With("key2", "value2").Fatalf("%s %s", "the", "message")
	if exitStatus != 1 {
		fmt.Printf("unexpected exit status: %d", exitStatus)
		return
	}
	fmt.Println(strings.Contains(writer.String(),
		`,"key":"value","key2":"value2","level":"fatal",`+machineName+`"msg":"the message",`+
			`"stack":"github.com/lovego/logger.ExampleFields_Fatalf\n\t`,
	))
	// Output: true
}
