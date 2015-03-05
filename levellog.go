package levellog

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Level int

const (
	DEBUG    Level = iota
	TRACE    Level = iota
	INFO     Level = iota
	WARN     Level = iota
	CRITICAL Level = iota
	PANIC    Level = iota
)

var (
	loggers      map[Level]*log.Logger
	currentLevel Level = DEBUG
	initCalled         = false
)

func SetLevel(s string) (err error) {
	l, err := ParseLevel(s) // returns DEBUG if error
	initLoggersOnce(l)
	return
}

func IsLevel(l Level) bool {
	return currentLevel >= l
}

func IsLevelString(s string) bool {
	l, _ := ParseLevel(s)
	return IsLevel(l)
}

// TODO: test init with writer interface
func initLoggers(l Level) {
	currentLevel = l
	loggers = make(map[Level]*log.Logger)

	for i := DEBUG; i <= PANIC; i++ {
		loggers[i] = log.New(os.Stderr, fmt.Sprintf("%s: ", i), log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func initLoggersOnce(l Level) {
	if initCalled {
		return
	}
	initLoggers(l)
	initCalled = true
}

func Printf(l Level, format string, v ...interface{}) {
	if l < currentLevel {
		return
	}
	loggers[l].Printf(format, v)
}

func Println(l Level, v ...interface{}) {
	if l < currentLevel {
		return
	}
	loggers[l].Println(v)
}

func Print(l Level, v ...interface{}) {
	if l < currentLevel {
		return
	}
	loggers[l].Print(v)
}

func Output(l Level, calldepth int, s string) error {
	return loggers[l].Output(calldepth, s)
}

func Panic(v ...interface{}) {
	loggers[PANIC].Panic(v)
}

func Panicf(format string, v ...interface{}) {
	loggers[PANIC].Panicf(format, v)
}

func Panicln(v ...interface{}) {
	loggers[PANIC].Panicln(v)
}

func Fatal(v ...interface{}) {
	loggers[CRITICAL].Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	loggers[CRITICAL].Fatalf(format, v)
}

func Fatalln(v ...interface{}) {
	loggers[CRITICAL].Fatalln(v)
}

// TODO: use stringer utility
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case CRITICAL:
		return "CRITICAL"
	case PANIC:
		return "PANIC"
	default:
		return "INVALID_LEVEL"
	}
}

func ParseLevel(s string) (Level, error) {
	strLevel := strings.TrimSpace(strings.ToUpper(s))
	switch strLevel {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARN":
		return WARN, nil
	case "CRITICAL":
		return CRITICAL, nil
	case "PANIC":
		return PANIC, nil
	default:
		return DEBUG, fmt.Errorf("invalid level \"%s\"", s)
	}
}
