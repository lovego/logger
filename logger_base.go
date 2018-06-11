package logger

func (l *Logger) output(
	level Level, msg string, fields map[string]interface{},
) {
	if fields == nil {
		fields = make(map[string]interface{})
	}
	for k, v := range l.fields {
		fields[k] = v
	}
	fields["at"] = Time.Now()
	fields["level"] = level
	fields["msg"] = msg

	l.writer.Write(l.formatter.Format(fields))
}

func (l *Logger) SetLevel(level Level) *logger {
	if level < Error {
		level = Error
	} else if level > Debug {
		level = Debug
	}
	l.level = level
	return l
}

func (l *Logger) SetAlarm(alarm Alarm) *logger {
	l.alarm = alarm
	return l
}

// Set a default field by key and value.
// Don't use "at", "level", "msg", they will be overwritten.
func (l *Logger) Set(key string, value interface{}) *Logger {
	l.fields[key] = value
	return l
}

// Set a default hostname field
func (l *Logger) SetHostname() *Logger {
	hostname, _ := os.Hostname()
	l.fields["hostname"] = hostname
	return l
}

// Set a default ip field
func (l *Logger) SetIP() *Logger {
	addrs, _ := net.InterfaceAddrs()
	slice := []string{}
	for _, addr := range addrs {
		ip := strings.Split(addr.String(), `/`)[0]
		IP := net.ParseIP(ip)
		if mask := IP.DefaultMask(); mask != nil && !IP.IsLoopback() {
			slice = append(slice, ip)
		}
	}
	l.fields["ip"] = slice
	return l
}

// Set a default pid field
func (l *Logger) SetPid() *Logger {
	l.Fields["pid"] = os.Getpid()
	return l
}
