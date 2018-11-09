package logger

import "fmt"

func ExampleJsonFormatter() {
	bytes, err := jsonFormatter.Format(map[string]interface{}{"query": "a=b&c=d"})
	fmt.Println(string(bytes), err)
	// Output:
	// {"query":"a=b&c=d"}
	//  <nil>
}

func ExampleReadableFormatter() {
	bytes, err := readableFormatter.Format(map[string]interface{}{
		"msg":   "some error",
		"stack": "the stack info",
		"query": "a=b&c=d",
	})
	fmt.Println(string(bytes), err)
	// Output:
	// some error
	// the stack info
	// {
	//   "query": "a=b&c=d"
	// }
	//  <nil>
}
