package logger

import (
	"fmt"
)

func ExampleSet() {
	logger := New(nil)
	logger.Set("key", "value")
	delete(logger.fields, "machineName")
	fmt.Println(logger.fields)
	// Output:
	// map[key:value]
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
