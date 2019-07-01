package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func ExampleSet() {
	logger := New(nil)
	logger.Set("key", "value")
	fmt.Println(logger.fields)
	// Output:
	// map[key:value]
}

func ExampleSetMachineName() {
	writer := bytes.NewBuffer(nil)
	logger := New(writer)
	logger.SetMachineName()
	if _, ok := logger.fields["machineName"]; !ok {
		fmt.Printf("unexpected logger %v", logger)
	}
	logger.Info("the message")

	hostname, _ := os.Hostname()
	expect := fmt.Sprintf(`"level":"info","machineName":"%s","msg":"the message"}
`, hostname)

	fmt.Println(strings.HasSuffix(writer.String(), expect))
	// Output:
	// true
}

func ExampleSetMachineIP() {
	logger := New(nil)
	logger.SetMachineIP()
	_, ok := logger.fields["machineIP"]
	fmt.Println(ok)
	// Output:
	// true
}

func ExampleSetPid() {
	logger := New(nil)
	logger.SetPid()
	_, ok := logger.fields["pid"]
	fmt.Println(ok)
	// Output:
	// true
}
