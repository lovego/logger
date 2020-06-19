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

func getStackField(skip int, args ...interface{}) map[string]interface{} {
	for _, arg := range args {
		if err, ok := arg.(interface {
			Error() string
			Stack() string
		}); ok {
			if stack := err.Stack(); stack != "" {
				return map[string]interface{}{"stack": stack}
			}
		}
	}
	if skip > 0 {
		return map[string]interface{}{"stack": errs.Stack(skip)}
	}
	return nil
}

func setStackField(fields map[string]interface{}, skip int, args ...interface{}) {
	for _, arg := range args {
		if err, ok := arg.(interface {
			Error() string
			Stack() string
		}); ok {
			if stack := err.Stack(); stack != "" {
				fields["stack"] = stack
				return
			}
		}
	}
	if skip > 0 {
		fields["stack"] = errs.Stack(skip)
	}
}
