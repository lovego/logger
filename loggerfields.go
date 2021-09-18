package logger

import (
	"net"
	"os"
	"strings"

	"github.com/lovego/errs"
)

// Set a default field by key and value.
// Don't use "level", "at", "msg", "stack", "duration" they will be overwritten.
func (l *Logger) Set(key string, value interface{}) *Logger {
	l.fields[key] = value
	return l
}

// Set a default ip field
func (l *Logger) SetMachineIP() *Logger {
	addrs, _ := net.InterfaceAddrs()
	slice := []string{}
	for _, addr := range addrs {
		ip := strings.Split(addr.String(), `/`)[0]
		IP := net.ParseIP(ip)
		if mask := IP.DefaultMask(); mask != nil && !IP.IsLoopback() {
			slice = append(slice, ip)
		}
	}
	l.fields["machineIP"] = slice
	return l
}

// Set a default pid field
func (l *Logger) SetPid() *Logger {
	l.fields["pid"] = os.Getpid()
	return l
}

func getExtraFields(args []interface{}, stack *errs.Stack) map[string]interface{} {
	fields := map[string]interface{}{}
	setExtraFields(fields, args, stack.IncrSkip())
	return fields
}

func setExtraFields(fields map[string]interface{}, args []interface{}, stack *errs.Stack) {
	if stackStr := getStack(args, stack.IncrSkip()); stackStr != "" {
		fields["stack"] = stackStr
	}
	if data := getData(args); data != nil {
		fields["data"] = data
	}
}

func getStack(args []interface{}, stack *errs.Stack) string {
	for _, arg := range args {
		if err, ok := arg.(interface {
			Error() string
			Stack() string
		}); ok {
			if stack := err.Stack(); stack != "" {
				return stack
			}
		}
	}
	return stack.IncrSkip().String()
}

func getData(args []interface{}) interface{} {
	for _, arg := range args {
		if err, ok := arg.(interface {
			Error() string
			Data() interface{}
		}); ok {
			if data := err.Data(); data != nil {
				return data
			}
		}
	}
	return nil
}
