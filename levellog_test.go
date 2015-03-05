package levellog

import (
	"flag"
	"testing"
)

var (
	userLevelString = flag.String("l", "debug", "log level")
)

func TestInit(t *testing.T) {
	initLoggers(TRACE)
	if currentLevel != TRACE {
		t.Fatalf("expected %s, found %s", TRACE, currentLevel)
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel("TRACE")
	if currentLevel != TRACE {
		t.Fatalf("expected %s, found %s", TRACE, currentLevel)
	}
}

func TestParseLevel(t *testing.T) {
	l := getUserLevel(t)
	t.Logf("user level is: %s", l)
}

func TestPrintfLevel(t *testing.T) {
	l := getUserLevel(t)
	initLoggers(WARN)
	Printf(l, "this is a message at level %s", l)
}

func TestPrintlnLevel(t *testing.T) {
	l := getUserLevel(t)
	initLoggers(INFO)
	Println(l, "do you see me?")
}

func TestIsLevel(t *testing.T) {
	SetLevel("warn")
	if !IsLevelString("trace") {
		t.Errorf("failed")
	}
}

func getUserLevel(t *testing.T) Level {
	flag.Parse()
	l, err := ParseLevel(*userLevelString)
	if err != nil {
		t.Errorf("error parsing level: %v", err)
	}
	return l
}
